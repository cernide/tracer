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
)

// BookmarkBuildReader is a Reader for the BookmarkBuild structure.
type BookmarkBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBookmarkBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBookmarkBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBookmarkBuildOK creates a BookmarkBuildOK with default headers values
func NewBookmarkBuildOK() *BookmarkBuildOK {
	return &BookmarkBuildOK{}
}

/*BookmarkBuildOK handles this case with default header values.

A successful response.
*/
type BookmarkBuildOK struct {
	Payload interface{}
}

func (o *BookmarkBuildOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/builds/{id}/bookmark][%d] bookmarkBuildOK  %+v", 200, o.Payload)
}

func (o *BookmarkBuildOK) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkBuildForbidden creates a BookmarkBuildForbidden with default headers values
func NewBookmarkBuildForbidden() *BookmarkBuildForbidden {
	return &BookmarkBuildForbidden{}
}

/*BookmarkBuildForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type BookmarkBuildForbidden struct {
	Payload interface{}
}

func (o *BookmarkBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/builds/{id}/bookmark][%d] bookmarkBuildForbidden  %+v", 403, o.Payload)
}

func (o *BookmarkBuildForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkBuildNotFound creates a BookmarkBuildNotFound with default headers values
func NewBookmarkBuildNotFound() *BookmarkBuildNotFound {
	return &BookmarkBuildNotFound{}
}

/*BookmarkBuildNotFound handles this case with default header values.

Resource does not exist.
*/
type BookmarkBuildNotFound struct {
	Payload string
}

func (o *BookmarkBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/builds/{id}/bookmark][%d] bookmarkBuildNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkBuildNotFound) GetPayload() string {
	return o.Payload
}

func (o *BookmarkBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
