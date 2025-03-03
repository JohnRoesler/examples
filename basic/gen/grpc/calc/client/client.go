// Code generated by goa v3.5.4, DO NOT EDIT.
//
// calc gRPC client
//
// Command:
// $ goa gen goa.design/examples/basic/design -o
// $(GOPATH)/src/goa.design/examples/basic

package client

import (
	"context"

	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli calcpb.CalcClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the calc service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: calcpb.NewCalcClient(cc),
		opts:    opts,
	}
}

// Add calls the "Add" function in calcpb.CalcClient interface.
func (c *Client) Add() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAddFunc(c.grpccli, c.opts...),
			EncodeAddRequest,
			DecodeAddResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
