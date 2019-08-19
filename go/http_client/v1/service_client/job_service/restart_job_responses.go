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

// RestartJobReader is a Reader for the RestartJob structure.
type RestartJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRestartJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestartJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRestartJobOK creates a RestartJobOK with default headers values
func NewRestartJobOK() *RestartJobOK {
	return &RestartJobOK{}
}

/*RestartJobOK handles this case with default header values.

A successful response.
*/
type RestartJobOK struct {
	Payload *service_model.V1Job
}

func (o *RestartJobOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/restart][%d] restartJobOK  %+v", 200, o.Payload)
}

func (o *RestartJobOK) GetPayload() *service_model.V1Job {
	return o.Payload
}

func (o *RestartJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartJobForbidden creates a RestartJobForbidden with default headers values
func NewRestartJobForbidden() *RestartJobForbidden {
	return &RestartJobForbidden{}
}

/*RestartJobForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RestartJobForbidden struct {
	Payload interface{}
}

func (o *RestartJobForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/restart][%d] restartJobForbidden  %+v", 403, o.Payload)
}

func (o *RestartJobForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RestartJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartJobNotFound creates a RestartJobNotFound with default headers values
func NewRestartJobNotFound() *RestartJobNotFound {
	return &RestartJobNotFound{}
}

/*RestartJobNotFound handles this case with default header values.

Resource does not exist.
*/
type RestartJobNotFound struct {
	Payload string
}

func (o *RestartJobNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/restart][%d] restartJobNotFound  %+v", 404, o.Payload)
}

func (o *RestartJobNotFound) GetPayload() string {
	return o.Payload
}

func (o *RestartJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
