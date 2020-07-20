// Code generated by go-swagger; DO NOT EDIT.

package project_service

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

// NewCreateTokenParams creates a new CreateTokenParams object
// with the default values initialized.
func NewCreateTokenParams() *CreateTokenParams {
	var ()
	return &CreateTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTokenParamsWithTimeout creates a new CreateTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTokenParamsWithTimeout(timeout time.Duration) *CreateTokenParams {
	var ()
	return &CreateTokenParams{

		timeout: timeout,
	}
}

// NewCreateTokenParamsWithContext creates a new CreateTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTokenParamsWithContext(ctx context.Context) *CreateTokenParams {
	var ()
	return &CreateTokenParams{

		Context: ctx,
	}
}

// NewCreateTokenParamsWithHTTPClient creates a new CreateTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTokenParamsWithHTTPClient(client *http.Client) *CreateTokenParams {
	var ()
	return &CreateTokenParams{
		HTTPClient: client,
	}
}

/*CreateTokenParams contains all the parameters to send to the API endpoint
for the create token operation typically these are written to a http.Request
*/
type CreateTokenParams struct {

	/*Body*/
	Body *models.ProjectProjectTokenCreateRequest
	/*Project*/
	Project string
	/*Role*/
	Role string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create token params
func (o *CreateTokenParams) WithTimeout(timeout time.Duration) *CreateTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create token params
func (o *CreateTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create token params
func (o *CreateTokenParams) WithContext(ctx context.Context) *CreateTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create token params
func (o *CreateTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) WithHTTPClient(client *http.Client) *CreateTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create token params
func (o *CreateTokenParams) WithBody(body *models.ProjectProjectTokenCreateRequest) *CreateTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create token params
func (o *CreateTokenParams) SetBody(body *models.ProjectProjectTokenCreateRequest) {
	o.Body = body
}

// WithProject adds the project to the create token params
func (o *CreateTokenParams) WithProject(project string) *CreateTokenParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create token params
func (o *CreateTokenParams) SetProject(project string) {
	o.Project = project
}

// WithRole adds the role to the create token params
func (o *CreateTokenParams) WithRole(role string) *CreateTokenParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the create token params
func (o *CreateTokenParams) SetRole(role string) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param role
	if err := r.SetPathParam("role", o.Role); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
