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
)

// ArchiveExperimentReader is a Reader for the ArchiveExperiment structure.
type ArchiveExperimentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveExperimentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveExperimentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewArchiveExperimentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArchiveExperimentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewArchiveExperimentOK creates a ArchiveExperimentOK with default headers values
func NewArchiveExperimentOK() *ArchiveExperimentOK {
	return &ArchiveExperimentOK{}
}

/*ArchiveExperimentOK handles this case with default header values.

A successful response.
*/
type ArchiveExperimentOK struct {
	Payload interface{}
}

func (o *ArchiveExperimentOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/experiments/{id}/archive][%d] archiveExperimentOK  %+v", 200, o.Payload)
}

func (o *ArchiveExperimentOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveExperimentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveExperimentForbidden creates a ArchiveExperimentForbidden with default headers values
func NewArchiveExperimentForbidden() *ArchiveExperimentForbidden {
	return &ArchiveExperimentForbidden{}
}

/*ArchiveExperimentForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ArchiveExperimentForbidden struct {
	Payload interface{}
}

func (o *ArchiveExperimentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/experiments/{id}/archive][%d] archiveExperimentForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveExperimentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveExperimentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveExperimentNotFound creates a ArchiveExperimentNotFound with default headers values
func NewArchiveExperimentNotFound() *ArchiveExperimentNotFound {
	return &ArchiveExperimentNotFound{}
}

/*ArchiveExperimentNotFound handles this case with default header values.

Resource does not exist.
*/
type ArchiveExperimentNotFound struct {
	Payload string
}

func (o *ArchiveExperimentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/experiments/{id}/archive][%d] archiveExperimentNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveExperimentNotFound) GetPayload() string {
	return o.Payload
}

func (o *ArchiveExperimentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
