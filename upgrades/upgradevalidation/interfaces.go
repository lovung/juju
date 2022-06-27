// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgradevalidation

import (
	"github.com/juju/names/v4"
	"github.com/juju/replicaset/v2"
	"github.com/juju/version/v2"

	"github.com/juju/juju/state"
)

// StatePool represents a point of use interface for getting the state from the
// pool.
type StatePool interface {
	MongoVersion() (string, error)
}

// State represents a point of use interface for modelling a current model.
type State interface {
	HasUpgradeSeriesLocks() (bool, error)
	Release() bool
	AllModelUUIDs() ([]string, error)
	MachineCountForSeries(series ...string) (int, error)
	MongoCurrentStatus() (*replicaset.Status, error)
	SetModelAgentVersion(newVersion version.Number, stream *string, ignoreAgentVersions bool) error
	AbortCurrentUpgrade() error
}

// Model defines a point of use interface for the model from state.
type Model interface {
	IsControllerModel() bool
	AgentVersion() (version.Number, error)
	Owner() names.UserTag
	Name() string
	MigrationMode() state.MigrationMode
}
