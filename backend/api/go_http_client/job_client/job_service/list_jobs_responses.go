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

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	job_model "github.com/kubeflow/pipelines/backend/api/go_http_client/job_model"
)

// ListJobsReader is a Reader for the ListJobs structure.
type ListJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListJobsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListJobsOK creates a ListJobsOK with default headers values
func NewListJobsOK() *ListJobsOK {
	return &ListJobsOK{}
}

/*ListJobsOK handles this case with default header values.

A successful response.
*/
type ListJobsOK struct {
	Payload *job_model.APIListJobsResponse
}

func (o *ListJobsOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/jobs][%d] listJobsOK  %+v", 200, o.Payload)
}

func (o *ListJobsOK) GetPayload() *job_model.APIListJobsResponse {
	return o.Payload
}

func (o *ListJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.APIListJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobsDefault creates a ListJobsDefault with default headers values
func NewListJobsDefault(code int) *ListJobsDefault {
	return &ListJobsDefault{
		_statusCode: code,
	}
}

/*ListJobsDefault handles this case with default header values.

ListJobsDefault list jobs default
*/
type ListJobsDefault struct {
	_statusCode int

	Payload *job_model.APIStatus
}

// Code gets the status code for the list jobs default response
func (o *ListJobsDefault) Code() int {
	return o._statusCode
}

func (o *ListJobsDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/jobs][%d] ListJobs default  %+v", o._statusCode, o.Payload)
}

func (o *ListJobsDefault) GetPayload() *job_model.APIStatus {
	return o.Payload
}

func (o *ListJobsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
