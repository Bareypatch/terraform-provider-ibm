// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudTenantsSshkeysGetParams creates a new PcloudTenantsSshkeysGetParams object
// with the default values initialized.
func NewPcloudTenantsSshkeysGetParams() *PcloudTenantsSshkeysGetParams {
	var ()
	return &PcloudTenantsSshkeysGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudTenantsSshkeysGetParamsWithTimeout creates a new PcloudTenantsSshkeysGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudTenantsSshkeysGetParamsWithTimeout(timeout time.Duration) *PcloudTenantsSshkeysGetParams {
	var ()
	return &PcloudTenantsSshkeysGetParams{

		timeout: timeout,
	}
}

// NewPcloudTenantsSshkeysGetParamsWithContext creates a new PcloudTenantsSshkeysGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudTenantsSshkeysGetParamsWithContext(ctx context.Context) *PcloudTenantsSshkeysGetParams {
	var ()
	return &PcloudTenantsSshkeysGetParams{

		Context: ctx,
	}
}

// NewPcloudTenantsSshkeysGetParamsWithHTTPClient creates a new PcloudTenantsSshkeysGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudTenantsSshkeysGetParamsWithHTTPClient(client *http.Client) *PcloudTenantsSshkeysGetParams {
	var ()
	return &PcloudTenantsSshkeysGetParams{
		HTTPClient: client,
	}
}

/*PcloudTenantsSshkeysGetParams contains all the parameters to send to the API endpoint
for the pcloud tenants sshkeys get operation typically these are written to a http.Request
*/
type PcloudTenantsSshkeysGetParams struct {

	/*SshkeyName
	  SSH key name for a pcloud tenant

	*/
	SshkeyName string
	/*TenantID
	  Tenant ID of a pcloud tenant

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) WithTimeout(timeout time.Duration) *PcloudTenantsSshkeysGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) WithContext(ctx context.Context) *PcloudTenantsSshkeysGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) WithHTTPClient(client *http.Client) *PcloudTenantsSshkeysGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSshkeyName adds the sshkeyName to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) WithSshkeyName(sshkeyName string) *PcloudTenantsSshkeysGetParams {
	o.SetSshkeyName(sshkeyName)
	return o
}

// SetSshkeyName adds the sshkeyName to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) SetSshkeyName(sshkeyName string) {
	o.SshkeyName = sshkeyName
}

// WithTenantID adds the tenantID to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) WithTenantID(tenantID string) *PcloudTenantsSshkeysGetParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the pcloud tenants sshkeys get params
func (o *PcloudTenantsSshkeysGetParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudTenantsSshkeysGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sshkey_name
	if err := r.SetPathParam("sshkey_name", o.SshkeyName); err != nil {
		return err
	}

	// path param tenant_id
	if err := r.SetPathParam("tenant_id", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
