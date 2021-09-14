// Code generated by goa v3.5.2, DO NOT EDIT.
//
// storage endpoints
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package storage

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "storage" service endpoints.
type Endpoints struct {
	List        goa.Endpoint
	Show        goa.Endpoint
	Add         goa.Endpoint
	Remove      goa.Endpoint
	Rate        goa.Endpoint
	MultiAdd    goa.Endpoint
	MultiUpdate goa.Endpoint
}

// NewEndpoints wraps the methods of the "storage" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:        NewListEndpoint(s),
		Show:        NewShowEndpoint(s),
		Add:         NewAddEndpoint(s),
		Remove:      NewRemoveEndpoint(s),
		Rate:        NewRateEndpoint(s),
		MultiAdd:    NewMultiAddEndpoint(s),
		MultiUpdate: NewMultiUpdateEndpoint(s),
	}
}

// Use applies the given middleware to all the "storage" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Add = m(e.Add)
	e.Remove = m(e.Remove)
	e.Rate = m(e.Rate)
	e.MultiAdd = m(e.MultiAdd)
	e.MultiUpdate = m(e.MultiUpdate)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "storage".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredBottleCollection(res, "tiny")
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "storage".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredBottle(res, view)
		return vres, nil
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "storage".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Bottle)
		return s.Add(ctx, p)
	}
}

// NewRemoveEndpoint returns an endpoint function that calls the method
// "remove" of service "storage".
func NewRemoveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemovePayload)
		return nil, s.Remove(ctx, p)
	}
}

// NewRateEndpoint returns an endpoint function that calls the method "rate" of
// service "storage".
func NewRateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(map[uint32][]string)
		return nil, s.Rate(ctx, p)
	}
}

// NewMultiAddEndpoint returns an endpoint function that calls the method
// "multi_add" of service "storage".
func NewMultiAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.([]*Bottle)
		return s.MultiAdd(ctx, p)
	}
}

// NewMultiUpdateEndpoint returns an endpoint function that calls the method
// "multi_update" of service "storage".
func NewMultiUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MultiUpdatePayload)
		return nil, s.MultiUpdate(ctx, p)
	}
}
