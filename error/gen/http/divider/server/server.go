// Code generated by goa v2.2.6, DO NOT EDIT.
//
// divider HTTP server
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package server

import (
	"context"
	"net/http"

	divider "goa.design/examples/error/gen/divider"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the divider service endpoint HTTP handlers.
type Server struct {
	Mounts        []*MountPoint
	IntegerDivide http.Handler
	Divide        http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the divider service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *divider.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"IntegerDivide", "GET", "/idiv/{a}/{b}"},
			{"Divide", "GET", "/div/{a}/{b}"},
		},
		IntegerDivide: NewIntegerDivideHandler(e.IntegerDivide, mux, decoder, encoder, errhandler, formatter),
		Divide:        NewDivideHandler(e.Divide, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "divider" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.IntegerDivide = m(s.IntegerDivide)
	s.Divide = m(s.Divide)
}

// Mount configures the mux to serve the divider endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountIntegerDivideHandler(mux, h.IntegerDivide)
	MountDivideHandler(mux, h.Divide)
}

// MountIntegerDivideHandler configures the mux to serve the "divider" service
// "integer_divide" endpoint.
func MountIntegerDivideHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/idiv/{a}/{b}", f)
}

// NewIntegerDivideHandler creates a HTTP handler which loads the HTTP request
// and calls the "divider" service "integer_divide" endpoint.
func NewIntegerDivideHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeIntegerDivideRequest(mux, decoder)
		encodeResponse = EncodeIntegerDivideResponse(encoder)
		encodeError    = EncodeIntegerDivideError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "integer_divide")
		ctx = context.WithValue(ctx, goa.ServiceKey, "divider")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDivideHandler configures the mux to serve the "divider" service
// "divide" endpoint.
func MountDivideHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/div/{a}/{b}", f)
}

// NewDivideHandler creates a HTTP handler which loads the HTTP request and
// calls the "divider" service "divide" endpoint.
func NewDivideHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDivideRequest(mux, decoder)
		encodeResponse = EncodeDivideResponse(encoder)
		encodeError    = EncodeDivideError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "divide")
		ctx = context.WithValue(ctx, goa.ServiceKey, "divider")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
