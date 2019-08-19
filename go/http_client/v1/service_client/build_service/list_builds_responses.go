// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package build_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// ListBuildsReader is a Reader for the ListBuilds structure.
type ListBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListBuildsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListBuildsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListBuildsOK creates a ListBuildsOK with default headers values
func NewListBuildsOK() *ListBuildsOK {
	return &ListBuildsOK{}
}

/*ListBuildsOK handles this case with default header values.

A successful response.
*/
type ListBuildsOK struct {
	Payload *service_model.V1ListBuildsResponse
}

func (o *ListBuildsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/builds][%d] listBuildsOK  %+v", 200, o.Payload)
}

func (o *ListBuildsOK) GetPayload() *service_model.V1ListBuildsResponse {
	return o.Payload
}

func (o *ListBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListBuildsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildsForbidden creates a ListBuildsForbidden with default headers values
func NewListBuildsForbidden() *ListBuildsForbidden {
	return &ListBuildsForbidden{}
}

/*ListBuildsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListBuildsForbidden struct {
	Payload interface{}
}

func (o *ListBuildsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/builds][%d] listBuildsForbidden  %+v", 403, o.Payload)
}

func (o *ListBuildsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListBuildsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildsNotFound creates a ListBuildsNotFound with default headers values
func NewListBuildsNotFound() *ListBuildsNotFound {
	return &ListBuildsNotFound{}
}

/*ListBuildsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListBuildsNotFound struct {
	Payload string
}

func (o *ListBuildsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/builds][%d] listBuildsNotFound  %+v", 404, o.Payload)
}

func (o *ListBuildsNotFound) GetPayload() string {
	return o.Payload
}

func (o *ListBuildsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
