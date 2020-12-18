// Code generated by goa v2.2.6, DO NOT EDIT.
//
// calc gRPC server
//
// Command:
// $ goa gen goa.design/examples/basic/design -o
// $(GOPATH)/src/goa.design/examples/basic

package server

import (
	"context"

	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
	"goa.design/goa"
	goagrpc "goa.design/goa/grpc"
)

// Server implements the calcpb.CalcServer interface.
type Server struct {
	AddH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the calc service endpoints.
func New(e *calc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		AddH: NewAddHandler(e.Add, uh),
	}
}

// NewAddHandler creates a gRPC handler which serves the "calc" service "add"
// endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in calcpb.CalcServer interface.
func (s *Server) Add(ctx context.Context, message *calcpb.AddRequest) (*calcpb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.AddResponse), nil
}
