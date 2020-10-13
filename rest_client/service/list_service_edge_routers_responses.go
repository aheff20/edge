// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// ListServiceEdgeRoutersReader is a Reader for the ListServiceEdgeRouters structure.
type ListServiceEdgeRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceEdgeRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceEdgeRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServiceEdgeRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServiceEdgeRoutersOK creates a ListServiceEdgeRoutersOK with default headers values
func NewListServiceEdgeRoutersOK() *ListServiceEdgeRoutersOK {
	return &ListServiceEdgeRoutersOK{}
}

/*ListServiceEdgeRoutersOK handles this case with default header values.

A list of edge routers
*/
type ListServiceEdgeRoutersOK struct {
	Payload *rest_model.ListEdgeRoutersEnvelope
}

func (o *ListServiceEdgeRoutersOK) Error() string {
	return fmt.Sprintf("[GET /services/{id}/edge-routers][%d] listServiceEdgeRoutersOK  %+v", 200, o.Payload)
}

func (o *ListServiceEdgeRoutersOK) GetPayload() *rest_model.ListEdgeRoutersEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEdgeRoutersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceEdgeRoutersUnauthorized creates a ListServiceEdgeRoutersUnauthorized with default headers values
func NewListServiceEdgeRoutersUnauthorized() *ListServiceEdgeRoutersUnauthorized {
	return &ListServiceEdgeRoutersUnauthorized{}
}

/*ListServiceEdgeRoutersUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListServiceEdgeRoutersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceEdgeRoutersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /services/{id}/edge-routers][%d] listServiceEdgeRoutersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListServiceEdgeRoutersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceEdgeRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
