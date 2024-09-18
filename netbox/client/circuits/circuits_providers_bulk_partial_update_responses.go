// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// CircuitsProvidersBulkPartialUpdateReader is a Reader for the CircuitsProvidersBulkPartialUpdate structure.
type CircuitsProvidersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersBulkPartialUpdateOK creates a CircuitsProvidersBulkPartialUpdateOK with default headers values
func NewCircuitsProvidersBulkPartialUpdateOK() *CircuitsProvidersBulkPartialUpdateOK {
	return &CircuitsProvidersBulkPartialUpdateOK{}
}

/*
CircuitsProvidersBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsProvidersBulkPartialUpdateOK circuits providers bulk partial update o k
*/
type CircuitsProvidersBulkPartialUpdateOK struct {
	Payload *models.Provider
}

// IsSuccess returns true when this circuits providers bulk partial update o k response has a 2xx status code
func (o *CircuitsProvidersBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers bulk partial update o k response has a 3xx status code
func (o *CircuitsProvidersBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers bulk partial update o k response has a 4xx status code
func (o *CircuitsProvidersBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers bulk partial update o k response has a 5xx status code
func (o *CircuitsProvidersBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers bulk partial update o k response a status code equal to that given
func (o *CircuitsProvidersBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits providers bulk partial update o k response
func (o *CircuitsProvidersBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *CircuitsProvidersBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/providers/][%d] circuitsProvidersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /circuits/providers/][%d] circuitsProvidersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersBulkPartialUpdateOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersBulkPartialUpdateDefault creates a CircuitsProvidersBulkPartialUpdateDefault with default headers values
func NewCircuitsProvidersBulkPartialUpdateDefault(code int) *CircuitsProvidersBulkPartialUpdateDefault {
	return &CircuitsProvidersBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsProvidersBulkPartialUpdateDefault describes a response with status code -1, with default header values.

CircuitsProvidersBulkPartialUpdateDefault circuits providers bulk partial update default
*/
type CircuitsProvidersBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits providers bulk partial update default response has a 2xx status code
func (o *CircuitsProvidersBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits providers bulk partial update default response has a 3xx status code
func (o *CircuitsProvidersBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits providers bulk partial update default response has a 4xx status code
func (o *CircuitsProvidersBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits providers bulk partial update default response has a 5xx status code
func (o *CircuitsProvidersBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits providers bulk partial update default response a status code equal to that given
func (o *CircuitsProvidersBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits providers bulk partial update default response
func (o *CircuitsProvidersBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsProvidersBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/providers/][%d] circuits_providers_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /circuits/providers/][%d] circuits_providers_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}