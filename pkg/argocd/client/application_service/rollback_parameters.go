// Code generated by go-swagger; DO NOT EDIT.

package application_service

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

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// NewRollbackParams creates a new RollbackParams object
// with the default values initialized.
func NewRollbackParams() *RollbackParams {
	var ()
	return &RollbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackParamsWithTimeout creates a new RollbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRollbackParamsWithTimeout(timeout time.Duration) *RollbackParams {
	var ()
	return &RollbackParams{

		timeout: timeout,
	}
}

// NewRollbackParamsWithContext creates a new RollbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewRollbackParamsWithContext(ctx context.Context) *RollbackParams {
	var ()
	return &RollbackParams{

		Context: ctx,
	}
}

// NewRollbackParamsWithHTTPClient creates a new RollbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRollbackParamsWithHTTPClient(client *http.Client) *RollbackParams {
	var ()
	return &RollbackParams{
		HTTPClient: client,
	}
}

/*RollbackParams contains all the parameters to send to the API endpoint
for the rollback operation typically these are written to a http.Request
*/
type RollbackParams struct {

	/*Body*/
	Body *models.ApplicationApplicationRollbackRequest
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rollback params
func (o *RollbackParams) WithTimeout(timeout time.Duration) *RollbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback params
func (o *RollbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback params
func (o *RollbackParams) WithContext(ctx context.Context) *RollbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback params
func (o *RollbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback params
func (o *RollbackParams) WithHTTPClient(client *http.Client) *RollbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback params
func (o *RollbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rollback params
func (o *RollbackParams) WithBody(body *models.ApplicationApplicationRollbackRequest) *RollbackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rollback params
func (o *RollbackParams) SetBody(body *models.ApplicationApplicationRollbackRequest) {
	o.Body = body
}

// WithName adds the name to the rollback params
func (o *RollbackParams) WithName(name string) *RollbackParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the rollback params
func (o *RollbackParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
