package test

import (
	"goa.design/examples/error/gen/divider"
	"goa.design/plugins/v3/testing"
)

// Endpoints wraps the "divider" service endpoints.
type Endpoints struct {
	IntegerDivide *testing.Endpoint
	Divide        *testing.Endpoint
}

// NewEndpoints wraps the "divider" service test endpoints.
func NewEndpoints(s divider.Service) *Endpoints {
	return &Endpoints{
		IntegerDivide: testing.NewEndpoint(divider.ServiceName, "IntegerDivide", divider.NewIntegerDivideEndpoint(s)),
		Divide:        testing.NewEndpoint(divider.ServiceName, "Divide", divider.NewDivideEndpoint(s)),
	}
}

// Endpoints returns the "divider" service endpoints wrapped with the testing
// middleware.
func (e *Endpoints) Endpoints() *divider.Endpoints {
	return &divider.Endpoints{
		IntegerDivide: e.IntegerDivide.Endpoint(),
		Divide:        e.Divide.Endpoint(),
	}
}
