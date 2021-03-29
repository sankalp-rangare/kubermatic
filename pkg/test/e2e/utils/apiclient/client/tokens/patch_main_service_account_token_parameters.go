// Code generated by go-swagger; DO NOT EDIT.

package tokens

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

// NewPatchMainServiceAccountTokenParams creates a new PatchMainServiceAccountTokenParams object
// with the default values initialized.
func NewPatchMainServiceAccountTokenParams() *PatchMainServiceAccountTokenParams {
	var ()
	return &PatchMainServiceAccountTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchMainServiceAccountTokenParamsWithTimeout creates a new PatchMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchMainServiceAccountTokenParamsWithTimeout(timeout time.Duration) *PatchMainServiceAccountTokenParams {
	var ()
	return &PatchMainServiceAccountTokenParams{

		timeout: timeout,
	}
}

// NewPatchMainServiceAccountTokenParamsWithContext creates a new PatchMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchMainServiceAccountTokenParamsWithContext(ctx context.Context) *PatchMainServiceAccountTokenParams {
	var ()
	return &PatchMainServiceAccountTokenParams{

		Context: ctx,
	}
}

// NewPatchMainServiceAccountTokenParamsWithHTTPClient creates a new PatchMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchMainServiceAccountTokenParamsWithHTTPClient(client *http.Client) *PatchMainServiceAccountTokenParams {
	var ()
	return &PatchMainServiceAccountTokenParams{
		HTTPClient: client,
	}
}

/*PatchMainServiceAccountTokenParams contains all the parameters to send to the API endpoint
for the patch main service account token operation typically these are written to a http.Request
*/
type PatchMainServiceAccountTokenParams struct {

	/*Body*/
	Body []uint8
	/*ServiceaccountID*/
	ServiceAccountID string
	/*TokenID*/
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) WithTimeout(timeout time.Duration) *PatchMainServiceAccountTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) WithContext(ctx context.Context) *PatchMainServiceAccountTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) WithHTTPClient(client *http.Client) *PatchMainServiceAccountTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) WithBody(body []uint8) *PatchMainServiceAccountTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) SetBody(body []uint8) {
	o.Body = body
}

// WithServiceAccountID adds the serviceaccountID to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) WithServiceAccountID(serviceaccountID string) *PatchMainServiceAccountTokenParams {
	o.SetServiceAccountID(serviceaccountID)
	return o
}

// SetServiceAccountID adds the serviceaccountId to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) SetServiceAccountID(serviceaccountID string) {
	o.ServiceAccountID = serviceaccountID
}

// WithTokenID adds the tokenID to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) WithTokenID(tokenID string) *PatchMainServiceAccountTokenParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the patch main service account token params
func (o *PatchMainServiceAccountTokenParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchMainServiceAccountTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param serviceaccount_id
	if err := r.SetPathParam("serviceaccount_id", o.ServiceAccountID); err != nil {
		return err
	}

	// path param token_id
	if err := r.SetPathParam("token_id", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
