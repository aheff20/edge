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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDetailPostureCheckParams creates a new DetailPostureCheckParams object
// with the default values initialized.
func NewDetailPostureCheckParams() *DetailPostureCheckParams {
	var ()
	return &DetailPostureCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetailPostureCheckParamsWithTimeout creates a new DetailPostureCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetailPostureCheckParamsWithTimeout(timeout time.Duration) *DetailPostureCheckParams {
	var ()
	return &DetailPostureCheckParams{

		timeout: timeout,
	}
}

// NewDetailPostureCheckParamsWithContext creates a new DetailPostureCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetailPostureCheckParamsWithContext(ctx context.Context) *DetailPostureCheckParams {
	var ()
	return &DetailPostureCheckParams{

		Context: ctx,
	}
}

// NewDetailPostureCheckParamsWithHTTPClient creates a new DetailPostureCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDetailPostureCheckParamsWithHTTPClient(client *http.Client) *DetailPostureCheckParams {
	var ()
	return &DetailPostureCheckParams{
		HTTPClient: client,
	}
}

/*DetailPostureCheckParams contains all the parameters to send to the API endpoint
for the detail posture check operation typically these are written to a http.Request
*/
type DetailPostureCheckParams struct {

	/*ID
	  The id of the requested resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detail posture check params
func (o *DetailPostureCheckParams) WithTimeout(timeout time.Duration) *DetailPostureCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detail posture check params
func (o *DetailPostureCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detail posture check params
func (o *DetailPostureCheckParams) WithContext(ctx context.Context) *DetailPostureCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detail posture check params
func (o *DetailPostureCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detail posture check params
func (o *DetailPostureCheckParams) WithHTTPClient(client *http.Client) *DetailPostureCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detail posture check params
func (o *DetailPostureCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the detail posture check params
func (o *DetailPostureCheckParams) WithID(id string) *DetailPostureCheckParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the detail posture check params
func (o *DetailPostureCheckParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DetailPostureCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
