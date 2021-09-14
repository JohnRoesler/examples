// Code generated by goa v3.5.2, DO NOT EDIT.
//
// text HTTP client types
//
// Command:
// $ goa gen goa.design/examples/encodings/text/design -o
// $(GOPATH)/src/goa.design/examples/encodings/text

package client

import (
	text "goa.design/examples/encodings/text/gen/text"
)

// NewConcatstringfieldMyConcatenationOK builds a "text" service
// "concatstringfield" endpoint result from a HTTP "OK" response.
func NewConcatstringfieldMyConcatenationOK(body string) *text.MyConcatenation {
	v := body
	res := &text.MyConcatenation{
		Stringfield: &v,
	}

	return res
}

// NewConcatbytesfieldMyConcatenationOK builds a "text" service
// "concatbytesfield" endpoint result from a HTTP "OK" response.
func NewConcatbytesfieldMyConcatenationOK(body []byte) *text.MyConcatenation {
	v := body
	res := &text.MyConcatenation{
		Bytesfield: v,
	}

	return res
}
