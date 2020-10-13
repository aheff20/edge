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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// PatchPostureCheckOKCode is the HTTP code returned for type PatchPostureCheckOK
const PatchPostureCheckOKCode int = 200

/*PatchPostureCheckOK The patch request was successful and the resource has been altered

swagger:response patchPostureCheckOK
*/
type PatchPostureCheckOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchPostureCheckOK creates PatchPostureCheckOK with default headers values
func NewPatchPostureCheckOK() *PatchPostureCheckOK {

	return &PatchPostureCheckOK{}
}

// WithPayload adds the payload to the patch posture check o k response
func (o *PatchPostureCheckOK) WithPayload(payload *rest_model.Empty) *PatchPostureCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch posture check o k response
func (o *PatchPostureCheckOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPostureCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPostureCheckBadRequestCode is the HTTP code returned for type PatchPostureCheckBadRequest
const PatchPostureCheckBadRequestCode int = 400

/*PatchPostureCheckBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchPostureCheckBadRequest
*/
type PatchPostureCheckBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchPostureCheckBadRequest creates PatchPostureCheckBadRequest with default headers values
func NewPatchPostureCheckBadRequest() *PatchPostureCheckBadRequest {

	return &PatchPostureCheckBadRequest{}
}

// WithPayload adds the payload to the patch posture check bad request response
func (o *PatchPostureCheckBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchPostureCheckBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch posture check bad request response
func (o *PatchPostureCheckBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPostureCheckBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPostureCheckUnauthorizedCode is the HTTP code returned for type PatchPostureCheckUnauthorized
const PatchPostureCheckUnauthorizedCode int = 401

/*PatchPostureCheckUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response patchPostureCheckUnauthorized
*/
type PatchPostureCheckUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchPostureCheckUnauthorized creates PatchPostureCheckUnauthorized with default headers values
func NewPatchPostureCheckUnauthorized() *PatchPostureCheckUnauthorized {

	return &PatchPostureCheckUnauthorized{}
}

// WithPayload adds the payload to the patch posture check unauthorized response
func (o *PatchPostureCheckUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchPostureCheckUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch posture check unauthorized response
func (o *PatchPostureCheckUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPostureCheckUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPostureCheckNotFoundCode is the HTTP code returned for type PatchPostureCheckNotFound
const PatchPostureCheckNotFoundCode int = 404

/*PatchPostureCheckNotFound The requested resource does not exist

swagger:response patchPostureCheckNotFound
*/
type PatchPostureCheckNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchPostureCheckNotFound creates PatchPostureCheckNotFound with default headers values
func NewPatchPostureCheckNotFound() *PatchPostureCheckNotFound {

	return &PatchPostureCheckNotFound{}
}

// WithPayload adds the payload to the patch posture check not found response
func (o *PatchPostureCheckNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchPostureCheckNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch posture check not found response
func (o *PatchPostureCheckNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPostureCheckNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
