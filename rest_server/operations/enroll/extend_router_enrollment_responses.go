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

package enroll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// ExtendRouterEnrollmentOKCode is the HTTP code returned for type ExtendRouterEnrollmentOK
const ExtendRouterEnrollmentOKCode int = 200

/*ExtendRouterEnrollmentOK A response containg the edge routers new signed certificates (server chain, server cert, CAs).

swagger:response extendRouterEnrollmentOK
*/
type ExtendRouterEnrollmentOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.EnrollmentCertsEnvelope `json:"body,omitempty"`
}

// NewExtendRouterEnrollmentOK creates ExtendRouterEnrollmentOK with default headers values
func NewExtendRouterEnrollmentOK() *ExtendRouterEnrollmentOK {

	return &ExtendRouterEnrollmentOK{}
}

// WithPayload adds the payload to the extend router enrollment o k response
func (o *ExtendRouterEnrollmentOK) WithPayload(payload *rest_model.EnrollmentCertsEnvelope) *ExtendRouterEnrollmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the extend router enrollment o k response
func (o *ExtendRouterEnrollmentOK) SetPayload(payload *rest_model.EnrollmentCertsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExtendRouterEnrollmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ExtendRouterEnrollmentUnauthorizedCode is the HTTP code returned for type ExtendRouterEnrollmentUnauthorized
const ExtendRouterEnrollmentUnauthorizedCode int = 401

/*ExtendRouterEnrollmentUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response extendRouterEnrollmentUnauthorized
*/
type ExtendRouterEnrollmentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewExtendRouterEnrollmentUnauthorized creates ExtendRouterEnrollmentUnauthorized with default headers values
func NewExtendRouterEnrollmentUnauthorized() *ExtendRouterEnrollmentUnauthorized {

	return &ExtendRouterEnrollmentUnauthorized{}
}

// WithPayload adds the payload to the extend router enrollment unauthorized response
func (o *ExtendRouterEnrollmentUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ExtendRouterEnrollmentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the extend router enrollment unauthorized response
func (o *ExtendRouterEnrollmentUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExtendRouterEnrollmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
