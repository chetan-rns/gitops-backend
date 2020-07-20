// Code generated by go-swagger; DO NOT EDIT.

package repository_service

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
	"github.com/go-openapi/swag"
)

// NewGetHelmChartsParams creates a new GetHelmChartsParams object
// with the default values initialized.
func NewGetHelmChartsParams() *GetHelmChartsParams {
	var ()
	return &GetHelmChartsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHelmChartsParamsWithTimeout creates a new GetHelmChartsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHelmChartsParamsWithTimeout(timeout time.Duration) *GetHelmChartsParams {
	var ()
	return &GetHelmChartsParams{

		timeout: timeout,
	}
}

// NewGetHelmChartsParamsWithContext creates a new GetHelmChartsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHelmChartsParamsWithContext(ctx context.Context) *GetHelmChartsParams {
	var ()
	return &GetHelmChartsParams{

		Context: ctx,
	}
}

// NewGetHelmChartsParamsWithHTTPClient creates a new GetHelmChartsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHelmChartsParamsWithHTTPClient(client *http.Client) *GetHelmChartsParams {
	var ()
	return &GetHelmChartsParams{
		HTTPClient: client,
	}
}

/*GetHelmChartsParams contains all the parameters to send to the API endpoint
for the get helm charts operation typically these are written to a http.Request
*/
type GetHelmChartsParams struct {

	/*ForceRefresh
	  Whether to force a cache refresh on repo's connection state.

	*/
	ForceRefresh *bool
	/*Repo*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get helm charts params
func (o *GetHelmChartsParams) WithTimeout(timeout time.Duration) *GetHelmChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get helm charts params
func (o *GetHelmChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get helm charts params
func (o *GetHelmChartsParams) WithContext(ctx context.Context) *GetHelmChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get helm charts params
func (o *GetHelmChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get helm charts params
func (o *GetHelmChartsParams) WithHTTPClient(client *http.Client) *GetHelmChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get helm charts params
func (o *GetHelmChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceRefresh adds the forceRefresh to the get helm charts params
func (o *GetHelmChartsParams) WithForceRefresh(forceRefresh *bool) *GetHelmChartsParams {
	o.SetForceRefresh(forceRefresh)
	return o
}

// SetForceRefresh adds the forceRefresh to the get helm charts params
func (o *GetHelmChartsParams) SetForceRefresh(forceRefresh *bool) {
	o.ForceRefresh = forceRefresh
}

// WithRepo adds the repo to the get helm charts params
func (o *GetHelmChartsParams) WithRepo(repo string) *GetHelmChartsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the get helm charts params
func (o *GetHelmChartsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *GetHelmChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceRefresh != nil {

		// query param forceRefresh
		var qrForceRefresh bool
		if o.ForceRefresh != nil {
			qrForceRefresh = *o.ForceRefresh
		}
		qForceRefresh := swag.FormatBool(qrForceRefresh)
		if qForceRefresh != "" {
			if err := r.SetQueryParam("forceRefresh", qForceRefresh); err != nil {
				return err
			}
		}

	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
