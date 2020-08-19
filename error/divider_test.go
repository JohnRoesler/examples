package divider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	dividersvc "goa.design/examples/error/gen/divider"
	"goa.design/examples/error/gen/divider/test"
	dividercli "goa.design/examples/error/gen/http/divider/client"
	dividersvr "goa.design/examples/error/gen/http/divider/server"
	goahttp "goa.design/goa/v3/http"
	"goa.design/goa/v3/http/middleware"
)

func TestIntegerDivide(t *testing.T) {
	svc := NewDivider(log.New(os.Stderr, "[dividerapi] ", log.Ltime))

	// create test endpoints for making assertions
	dividerTest := test.NewEndpoints(svc)

	// Create http test server for handling HTTP requests to divider service
	var mux = goahttp.NewMuxer()
	{
		var httpService = dividersvr.New(dividerTest.Endpoints(), mux,
			goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		httpService.Use(middleware.PopulateRequestContext())
		dividersvr.Mount(mux, httpService)
	}

	httpServer := httptest.NewServer(mux)

	var (
		httpServerURL *url.URL
		err           error
	)
	{
		httpServerURL, err = url.Parse(httpServer.URL)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	// create front service client to make requests to test HTTP server
	httpClient := dividercli.NewClient(httpServerURL.Scheme, httpServerURL.Host,
		&http.Client{}, goahttp.RequestEncoder, goahttp.ResponseDecoder, false)

	t.Run("success", func(t *testing.T) {
		payload := &dividersvc.IntOperands{A: 4, B: 2}
		result := 2

		// set expectations on the endpoint
		dividerTest.IntegerDivide.WillGet(payload).WillReturn(result)

		// make HTTP request
		res, err := httpClient.IntegerDivide()(context.Background(), payload)
		if err != nil {
			t.Errorf(err.Error())
		}
		if res.(int) != result {
			t.Errorf("expected %d, got %d", result, res)
		}
	})

	t.Run("divide by zero error", func(t *testing.T) {
		payload := &dividersvc.IntOperands{A: 4, B: 0}

		// set expectations on the endpoint
		dividerTest.IntegerDivide.WillGet(payload).WillReturnError(dividersvc.MakeDivByZero(fmt.Errorf("div by 0")))

		// make HTTP request
		_, err := httpClient.IntegerDivide()(context.Background(), payload)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
