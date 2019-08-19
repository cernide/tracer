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

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// CreateJobStatusReader is a Reader for the CreateJobStatus structure.
type CreateJobStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateJobStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateJobStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateJobStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateJobStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateJobStatusOK creates a CreateJobStatusOK with default headers values
func NewCreateJobStatusOK() *CreateJobStatusOK {
	return &CreateJobStatusOK{}
}

/*CreateJobStatusOK handles this case with default header values.

A successful response.
*/
type CreateJobStatusOK struct {
	Payload *service_model.V1JobStatus
}

func (o *CreateJobStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/statuses][%d] createJobStatusOK  %+v", 200, o.Payload)
}

func (o *CreateJobStatusOK) GetPayload() *service_model.V1JobStatus {
	return o.Payload
}

func (o *CreateJobStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobStatusForbidden creates a CreateJobStatusForbidden with default headers values
func NewCreateJobStatusForbidden() *CreateJobStatusForbidden {
	return &CreateJobStatusForbidden{}
}

/*CreateJobStatusForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateJobStatusForbidden struct {
	Payload interface{}
}

func (o *CreateJobStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/statuses][%d] createJobStatusForbidden  %+v", 403, o.Payload)
}

func (o *CreateJobStatusForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateJobStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobStatusNotFound creates a CreateJobStatusNotFound with default headers values
func NewCreateJobStatusNotFound() *CreateJobStatusNotFound {
	return &CreateJobStatusNotFound{}
}

/*CreateJobStatusNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateJobStatusNotFound struct {
	Payload string
}

func (o *CreateJobStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/statuses][%d] createJobStatusNotFound  %+v", 404, o.Payload)
}

func (o *CreateJobStatusNotFound) GetPayload() string {
	return o.Payload
}

func (o *CreateJobStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
