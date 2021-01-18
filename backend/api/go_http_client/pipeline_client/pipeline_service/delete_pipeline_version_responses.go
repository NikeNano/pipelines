// Copyright 2021 Google LLC
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

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubeflow/pipelines/backend/api/go_http_client/pipeline_model"
)

// DeletePipelineVersionReader is a Reader for the DeletePipelineVersion structure.
type DeletePipelineVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePipelineVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePipelineVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePipelineVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePipelineVersionOK creates a DeletePipelineVersionOK with default headers values
func NewDeletePipelineVersionOK() *DeletePipelineVersionOK {
	return &DeletePipelineVersionOK{}
}

/*DeletePipelineVersionOK handles this case with default header values.

A successful response.
*/
type DeletePipelineVersionOK struct {
	Payload interface{}
}

func (o *DeletePipelineVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/v1beta1/pipeline_versions/{version_id}][%d] deletePipelineVersionOK  %+v", 200, o.Payload)
}

func (o *DeletePipelineVersionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeletePipelineVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePipelineVersionDefault creates a DeletePipelineVersionDefault with default headers values
func NewDeletePipelineVersionDefault(code int) *DeletePipelineVersionDefault {
	return &DeletePipelineVersionDefault{
		_statusCode: code,
	}
}

/*DeletePipelineVersionDefault handles this case with default header values.

DeletePipelineVersionDefault delete pipeline version default
*/
type DeletePipelineVersionDefault struct {
	_statusCode int

	Payload *pipeline_model.APIStatus
}

// Code gets the status code for the delete pipeline version default response
func (o *DeletePipelineVersionDefault) Code() int {
	return o._statusCode
}

func (o *DeletePipelineVersionDefault) Error() string {
	return fmt.Sprintf("[DELETE /apis/v1beta1/pipeline_versions/{version_id}][%d] DeletePipelineVersion default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePipelineVersionDefault) GetPayload() *pipeline_model.APIStatus {
	return o.Payload
}

func (o *DeletePipelineVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
