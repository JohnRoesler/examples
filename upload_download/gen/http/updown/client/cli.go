// Code generated by goa v3.5.4, DO NOT EDIT.
//
// updown HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/upload_download/design -o
// $(GOPATH)/src/goa.design/examples/upload_download

package client

import (
	updown "goa.design/examples/upload_download/gen/updown"
	goa "goa.design/goa/v3/pkg"
)

// BuildUploadPayload builds the payload for the updown upload endpoint from
// CLI flags.
func BuildUploadPayload(updownUploadDir string, updownUploadContentType string) (*updown.UploadPayload, error) {
	var err error
	var dir string
	{
		dir = updownUploadDir
	}
	var contentType string
	{
		if updownUploadContentType != "" {
			contentType = updownUploadContentType
			err = goa.MergeErrors(err, goa.ValidatePattern("contentType", contentType, "multipart/[^;]+; boundary=.+"))
			if err != nil {
				return nil, err
			}
		}
	}
	v := &updown.UploadPayload{}
	v.Dir = dir
	v.ContentType = contentType

	return v, nil
}
