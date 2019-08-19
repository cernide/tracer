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

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// ListBookmarkedExperimentsReader is a Reader for the ListBookmarkedExperiments structure.
type ListBookmarkedExperimentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBookmarkedExperimentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBookmarkedExperimentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListBookmarkedExperimentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListBookmarkedExperimentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListBookmarkedExperimentsOK creates a ListBookmarkedExperimentsOK with default headers values
func NewListBookmarkedExperimentsOK() *ListBookmarkedExperimentsOK {
	return &ListBookmarkedExperimentsOK{}
}

/*ListBookmarkedExperimentsOK handles this case with default header values.

A successful response.
*/
type ListBookmarkedExperimentsOK struct {
	Payload *service_model.V1ListExperimentsResponse
}

func (o *ListBookmarkedExperimentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{owner}/experiments][%d] listBookmarkedExperimentsOK  %+v", 200, o.Payload)
}

func (o *ListBookmarkedExperimentsOK) GetPayload() *service_model.V1ListExperimentsResponse {
	return o.Payload
}

func (o *ListBookmarkedExperimentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListExperimentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBookmarkedExperimentsForbidden creates a ListBookmarkedExperimentsForbidden with default headers values
func NewListBookmarkedExperimentsForbidden() *ListBookmarkedExperimentsForbidden {
	return &ListBookmarkedExperimentsForbidden{}
}

/*ListBookmarkedExperimentsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListBookmarkedExperimentsForbidden struct {
	Payload interface{}
}

func (o *ListBookmarkedExperimentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{owner}/experiments][%d] listBookmarkedExperimentsForbidden  %+v", 403, o.Payload)
}

func (o *ListBookmarkedExperimentsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListBookmarkedExperimentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBookmarkedExperimentsNotFound creates a ListBookmarkedExperimentsNotFound with default headers values
func NewListBookmarkedExperimentsNotFound() *ListBookmarkedExperimentsNotFound {
	return &ListBookmarkedExperimentsNotFound{}
}

/*ListBookmarkedExperimentsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListBookmarkedExperimentsNotFound struct {
	Payload string
}

func (o *ListBookmarkedExperimentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/bookmarks/{owner}/experiments][%d] listBookmarkedExperimentsNotFound  %+v", 404, o.Payload)
}

func (o *ListBookmarkedExperimentsNotFound) GetPayload() string {
	return o.Payload
}

func (o *ListBookmarkedExperimentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
