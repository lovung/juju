// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgradesteps

import (
	"reflect"

	"github.com/juju/juju/apiserver/facade"
)

// Registry describes the API facades exposed by some API server.
type Registry interface {
	// MustRegister adds a single named facade at a given version to the
	// registry.
	// Factory will be called when someone wants to instantiate an object of
	// this facade, and facadeType defines the concrete type that the returned
	// object will be.
	// The Type information is used to define what methods will be exported in
	// the API, and it must exactly match the actual object returned by the
	// factory.
	MustRegister(string, int, facade.Factory, reflect.Type)
}

// Register is called to expose a package of facades onto a given registry.
func Register(registry Registry) {
	registry.MustRegister("UpgradeSteps", 1, func(ctx facade.Context) (facade.Facade, error) {
		return newFacadeV1(ctx)
	}, reflect.TypeOf((*UpgradeStepsAPIV1)(nil)))
	registry.MustRegister("UpgradeSteps", 2, func(ctx facade.Context) (facade.Facade, error) {
		return newFacadeV2(ctx)
	}, reflect.TypeOf((*UpgradeStepsAPI)(nil)))
}

// newFacadeV2 is used for API registration.
func newFacadeV2(ctx facade.Context) (*UpgradeStepsAPI, error) {
	st := &upgradeStepsStateShim{State: ctx.State()}
	return NewUpgradeStepsAPI(st, ctx.Resources(), ctx.Auth())
}

// newFacadeV1 is used for API registration.
func newFacadeV1(ctx facade.Context) (*UpgradeStepsAPIV1, error) {
	v2, err := newFacadeV2(ctx)
	if err != nil {
		return nil, err
	}
	return &UpgradeStepsAPIV1{UpgradeStepsAPI: v2}, nil
}
