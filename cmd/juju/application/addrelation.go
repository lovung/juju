// Copyright 2012, 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package application

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/juju/cmd/v3"
	"github.com/juju/errors"
	"github.com/juju/gnuflag"
	"github.com/juju/names/v4"

	"github.com/juju/juju/api/client/application"
	"github.com/juju/juju/api/client/applicationoffers"
	jujucmd "github.com/juju/juju/cmd"
	"github.com/juju/juju/cmd/juju/block"
	"github.com/juju/juju/cmd/juju/common"
	"github.com/juju/juju/cmd/modelcmd"
	"github.com/juju/juju/core/crossmodel"
	"github.com/juju/juju/rpc/params"
)

const addRelationDoc = `
Relate two applications. Relations are a communication interface provided by 
the Juju controller that enable units to transfer data. This star messaging 
topology allows units to send and receive data, even if direct connectivity
between units is restricted by firewall rules. Charms define the logic for 
transferring and interpreting relation data.

The most common use of 'juju relate' specifies two applications that co-exist
within the same model:

    juju relate <application> <application> 

Occasionally, more explicit syntax is required. Juju is able to relate 
units that span models, controllers and clouds, as described below.


Relating applications in the same model

The most common case specifying two applications, and adding the specific 
relation name when required.

    juju relate <application>[:<relation>] <application>[:<relation>]

Juju uses these arguments to create an "application endpoint". An application
endpoint is the combination of an "application name", a "relation name" and a 
"role". The role and relation name are described by charms' metadata.yaml file.

The order does not matter, however each side must perform complementary roles.
One side performs the "provides" role and the other side performs the "requires" 
role. Juju can always infer the role that each side is performing, so specifying
them is not necessary as command-line arguments.

<app-name> is the name of an application that has already been added to the 
model. The Applications section of 'juju status' provides a list of current
applications.

<relation-name> is the name of a relation defined within the metadata.yaml 
of the charm for <app-name>. Valid relation names are defined within the 
"provides:" and "requires:" section of that file. Juju will request that you
specify the <relation-name> when it is unable to resolve the name itself.


Subordinate applications

Relating a principal application to a subordinate application has the effect of
deploying the subordinate alongside its principal. This functionality use the 
same syntax as relating applications within the same model.
 

Peer relations

Relations within an application between units (known as "peer relations") do 
not need to be added manually. They are created when the 'juju add-unit' and
'juju scale-application' commands are executed. 


Cross-model relations

Applications can be related, even when they are deployed to different models.
Those models may be managed by different controllers and/or be hosted on 
different clouds. This functionality is known as "cross-model relations" or CMR.


Cross-model relations: different model on the same controller

Adding a relation between applications in models managed by the same controller
is very similar to adding a relation between applications in the same model:

    juju relate <application>[:<relation>] <model>.<application>[:<relation>]

<model> is the name of the model outside of the current context. enables the 
Juju controller to bridge two models. List the model currently available with
the 'juju models' command.

To relate models outside of the current context, add the '-m <model>' option:

    juju relate -m <model> <application>[:<relation>] \
                   <model>.<application>[:<relation>]


Cross-model relations: different controllers

Applications can relate to a remove application via an "offer URL" that has 
been generated by the 'juju offer' command. The syntax for adding a cross-model 
relation is similar to adding a local relation:

    juju relate <application>[:<relation>] <offer-endpoint>

<offer-endpoint> describes the remote application, from the point of view of the
local one. An <offer-endpoint> takes one of two forms:

    <offer-alias>
    <offer-url>[:<relation-name>]

<offer-alias> is an alias that has been defined by the 'juju consume' command. 
Use the 'juju find-offers' command to list aliases.

<offer-url> is a path to enable Juju to resolve communication between 
controllers and the models they control.

    [[<controller>:]<user>/]<model-name>.<application-name>

<controller> is the name of a controller. The 'juju controllers' command 
provides a list of controllers.

<user> is the user account of the model's owner. 


Cross-model relations: network management

When the consuming side (the local application) is behind a firewall and/or 
NAT is used for outbound traffic, it is possible to use the '--via' option to 
inform the offering side (the remote application) the source of traffic to 
enable network ports to be opened.

    ... --via <cidr-subnet>[,<cidr-subnet>[, ...]] 


Further reading:

    https://juju.is/docs/relations
    https://juju.is/docs/cross-model-relations


Examples:
    
    # Relate the wordpress and percona-cluster applications, asking Juju to resolve
    # the relation names. Expands to "wordpress:db" (with the requires role) and 
    # "percona-cluster:server" (with the provides role). 
    juju relate wordpress percona-cluster

    # Relate the wordpress and postgresql applications, using an explicit
    # relation name.
    juju relate wordpress postgresql:db

    # Relate an etcd instance within the current model to centrally managed
    # EasyRSA Certificate Authority hosted in the "secrets" model
    juju relate etcd secrets.easyrsa

    # Relate a wordpress application with a mysql application hosted within the 
    # "prod" model, using the "automation" user. Facilitate firewall management
    # by specifying the routes used for relation data.
    juju add-relation wordpress automation/prod.mysql --via 192.168.0.0/16,10.0.0.0/8


See also:

    consume
    find-offers
    set-firewall-rule
    suspend-relation
`

var localEndpointRegEx = regexp.MustCompile("^" + names.RelationSnippet + "$")

// NewAddRelationCommand returns a command to add a relation between 2 applications.
func NewAddRelationCommand() cmd.Command {
	return modelcmd.Wrap(&addRelationCommand{})
}

// addRelationCommand adds a relation between two application endpoints.
type addRelationCommand struct {
	modelcmd.ModelCommandBase
	endpoints         []string
	viaCIDRs          []string
	viaValue          string
	remoteEndpoint    *crossmodel.OfferURL
	addRelationAPI    applicationAddRelationAPI
	consumeDetailsAPI applicationConsumeDetailsAPI
}

func (c *addRelationCommand) Info() *cmd.Info {
	addCmd := &cmd.Info{
		Name:    "add-relation",
		Aliases: []string{"relate"},
		Args:    "<application>[:<relation>] <application>[:<relation>]",
		Purpose: "Relate two applications.",
		Doc:     addRelationDoc,
	}
	return jujucmd.Info(addCmd)
}

func (c *addRelationCommand) Init(args []string) error {
	if len(args) != 2 {
		return errors.Errorf("a relation must involve two applications")
	}
	if err := c.validateEndpoints(args); err != nil {
		return err
	}
	if err := c.validateCIDRs(); err != nil {
		return err
	}
	if c.remoteEndpoint == nil && len(c.viaCIDRs) > 0 {
		return errors.New("the --via option can only be used when relating to offers in a different model")
	}
	return nil
}

func (c *addRelationCommand) SetFlags(f *gnuflag.FlagSet) {
	f.StringVar(&c.viaValue, "via", "", "for cross model relations, specify the egress subnets for outbound traffic")
}

// applicationAddRelationAPI defines the API methods that application add relation command uses.
type applicationAddRelationAPI interface {
	Close() error
	BestAPIVersion() int
	AddRelation(endpoints, viaCIDRs []string) (*params.AddRelationResults, error)
	Consume(crossmodel.ConsumeApplicationArgs) (string, error)
}

func (c *addRelationCommand) getAddRelationAPI() (applicationAddRelationAPI, error) {
	if c.addRelationAPI != nil {
		return c.addRelationAPI, nil
	}

	root, err := c.NewAPIRoot()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return application.NewClient(root), nil
}

func (c *addRelationCommand) getOffersAPI(url *crossmodel.OfferURL) (applicationConsumeDetailsAPI, error) {
	if c.consumeDetailsAPI != nil {
		return c.consumeDetailsAPI, nil
	}

	root, err := c.CommandBase.NewAPIRoot(c.ClientStore(), url.Source, "")
	if err != nil {
		return nil, errors.Trace(err)
	}
	return applicationoffers.NewClient(root), nil
}

// offerTerminatedRegexp is used to parse an error due to the remote offer being terminated.
// (TODO) we don't have an error code for this scenario so need to rely on a string match.
var offerTerminatedRegexp = regexp.MustCompile(".*offer (?P<offer>\\S+) .*terminated.*")

func (c *addRelationCommand) Run(ctx *cmd.Context) error {
	client, err := c.getAddRelationAPI()
	if err != nil {
		return err
	}
	defer client.Close()

	if c.remoteEndpoint != nil {
		if client.BestAPIVersion() < 5 {
			// old client does not have cross-model capability.
			return errors.NotSupportedf("cannot add relation to %s: remote endpoints", c.remoteEndpoint.String())
		}
		if c.remoteEndpoint.Source == "" {
			var err error
			controllerName, err := c.ControllerName()
			if err != nil {
				return errors.Trace(err)
			}
			c.remoteEndpoint.Source = controllerName
		}
		if err := c.maybeConsumeOffer(client); err != nil {
			return errors.Trace(err)
		}
	}

	_, err = client.AddRelation(c.endpoints, c.viaCIDRs)
	if params.IsCodeUnauthorized(err) {
		common.PermissionsMessage(ctx.Stderr, "add a relation")
	}
	if params.IsCodeAlreadyExists(err) {
		splitError := strings.Join(strings.Split(err.Error(), ": "), "\n")
		infoErr := errors.Errorf(`

Use 'juju status --relations' to view the current relations.`)
		return errors.Annotatef(infoErr, splitError)
	}
	if err != nil {
		if offerTerminatedRegexp.MatchString(err.Error()) {
			offerName := offerTerminatedRegexp.ReplaceAllString(err.Error(), "$offer")
			return errors.New(fmt.Sprintf(
				`Offer %q has been removed from the remote model.
To relate to a new offer with the same name, first run
'juju remove-saas %s' to remove the SAAS record from this model.`, offerName, offerName))
		}
		if c.remoteEndpoint != nil && strings.HasSuffix(err.Error(), "not alive") {
			saasName := c.remoteEndpoint.ApplicationName
			return errors.New(fmt.Sprintf(
				`SAAS application %q has been removed but termination has not completed.
To relate to a new offer with the same name, first run
'juju remove-saas %s --force' to remove the SAAS record from this model.`, saasName, saasName))
		}
	}
	return block.ProcessBlockedError(err, block.BlockChange)
}

func (c *addRelationCommand) maybeConsumeOffer(targetClient applicationAddRelationAPI) error {
	sourceClient, err := c.getOffersAPI(c.remoteEndpoint)
	if err != nil {
		return errors.Trace(err)
	}
	defer sourceClient.Close()

	// Get the details of the remote offer - this will fail with a permission
	// error if the user isn't authorised to consume the offer.
	consumeDetails, err := sourceClient.GetConsumeDetails(c.remoteEndpoint.AsLocal().String())
	if err != nil {
		return errors.Trace(err)
	}
	// Parse the offer details URL and add the source controller so
	// things like status can show the original source of the offer.
	offerURL, err := crossmodel.ParseOfferURL(consumeDetails.Offer.OfferURL)
	if err != nil {
		return errors.Trace(err)
	}
	offerURL.Source = c.remoteEndpoint.Source
	consumeDetails.Offer.OfferURL = offerURL.String()

	// Consume is idempotent so even if the offer has been consumed previously,
	// it's safe to do so again.
	arg := crossmodel.ConsumeApplicationArgs{
		Offer:            *consumeDetails.Offer,
		ApplicationAlias: c.remoteEndpoint.ApplicationName,
		Macaroon:         consumeDetails.Macaroon,
	}
	if consumeDetails.ControllerInfo != nil {
		controllerTag, err := names.ParseControllerTag(consumeDetails.ControllerInfo.ControllerTag)
		if err != nil {
			return errors.Trace(err)
		}
		arg.ControllerInfo = &crossmodel.ControllerInfo{
			ControllerTag: controllerTag,
			Alias:         offerURL.Source,
			Addrs:         consumeDetails.ControllerInfo.Addrs,
			CACert:        consumeDetails.ControllerInfo.CACert,
		}
	}
	_, err = targetClient.Consume(arg)
	return errors.Trace(err)
}

// validateEndpoints determines if all endpoints are valid.
// Each endpoint is either from local application or remote.
// If more than one remote endpoint are supplied, the input argument are considered invalid.
func (c *addRelationCommand) validateEndpoints(all []string) error {
	for _, endpoint := range all {
		// We can only determine if this is a remote endpoint with 100%.
		// If we cannot parse it, it may still be a valid local endpoint...
		// so ignoring parsing error,
		if url, err := crossmodel.ParseOfferURL(endpoint); err == nil {
			if c.remoteEndpoint != nil {
				return errors.NotSupportedf("providing more than one remote endpoints")
			}
			c.remoteEndpoint = url
			c.endpoints = append(c.endpoints, url.ApplicationName)
			continue
		}
		// at this stage, we are assuming that this could be a local endpoint
		if err := validateLocalEndpoint(endpoint, ":"); err != nil {
			return err
		}
		c.endpoints = append(c.endpoints, endpoint)
	}
	return nil
}

// validateLocalEndpoint determines if given endpoint could be a valid
func validateLocalEndpoint(endpoint string, sep string) error {
	i := strings.Index(endpoint, sep)
	applicationName := endpoint
	if i != -1 {
		// not a valid endpoint as sep either at the start or the end of the name
		if i == 0 || i == len(endpoint)-1 {
			return errors.NotValidf("endpoint %q", endpoint)
		}

		parts := strings.SplitN(endpoint, sep, -1)
		if rightCount := len(parts) == 2; !rightCount {
			// not valid if there are not exactly 2 parts.
			return errors.NotValidf("endpoint %q", endpoint)
		}

		applicationName = parts[0]

		if valid := localEndpointRegEx.MatchString(parts[1]); !valid {
			return errors.NotValidf("endpoint %q", endpoint)
		}
	}

	if valid := names.IsValidApplication(applicationName); !valid {
		return errors.NotValidf("application name %q", applicationName)
	}
	return nil
}

func (c *addRelationCommand) validateCIDRs() error {
	if c.viaValue == "" {
		return nil
	}
	c.viaCIDRs = strings.Split(
		strings.Replace(c.viaValue, " ", "", -1), ",")
	for _, cidr := range c.viaCIDRs {
		if _, _, err := net.ParseCIDR(cidr); err != nil {
			return err
		}
		if cidr == "0.0.0.0/0" {
			return errors.Errorf("CIDR %q not allowed", cidr)
		}
	}
	return nil
}
