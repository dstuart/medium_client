// Code generated by go-swagger; DO NOT EDIT.

package posts

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

	"github.com/dstuart/models"
)

// NewPostPublicationsPublicationIDPostsParams creates a new PostPublicationsPublicationIDPostsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPublicationsPublicationIDPostsParams() *PostPublicationsPublicationIDPostsParams {
	return &PostPublicationsPublicationIDPostsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicationsPublicationIDPostsParamsWithTimeout creates a new PostPublicationsPublicationIDPostsParams object
// with the ability to set a timeout on a request.
func NewPostPublicationsPublicationIDPostsParamsWithTimeout(timeout time.Duration) *PostPublicationsPublicationIDPostsParams {
	return &PostPublicationsPublicationIDPostsParams{
		timeout: timeout,
	}
}

// NewPostPublicationsPublicationIDPostsParamsWithContext creates a new PostPublicationsPublicationIDPostsParams object
// with the ability to set a context for a request.
func NewPostPublicationsPublicationIDPostsParamsWithContext(ctx context.Context) *PostPublicationsPublicationIDPostsParams {
	return &PostPublicationsPublicationIDPostsParams{
		Context: ctx,
	}
}

// NewPostPublicationsPublicationIDPostsParamsWithHTTPClient creates a new PostPublicationsPublicationIDPostsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPublicationsPublicationIDPostsParamsWithHTTPClient(client *http.Client) *PostPublicationsPublicationIDPostsParams {
	return &PostPublicationsPublicationIDPostsParams{
		HTTPClient: client,
	}
}

/*
PostPublicationsPublicationIDPostsParams contains all the parameters to send to the API endpoint

	for the post publications publication ID posts operation.

	Typically these are written to a http.Request.
*/
type PostPublicationsPublicationIDPostsParams struct {

	/* Body.

	   Creates a post for publication.
	*/
	Body *models.Post

	/* PublicationID.

	   Here publicationId is the id of the publication the post is being created under. The publicationId can be acquired from the API for listing user’s publications.
	*/
	PublicationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post publications publication ID posts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPublicationsPublicationIDPostsParams) WithDefaults() *PostPublicationsPublicationIDPostsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post publications publication ID posts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPublicationsPublicationIDPostsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) WithTimeout(timeout time.Duration) *PostPublicationsPublicationIDPostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) WithContext(ctx context.Context) *PostPublicationsPublicationIDPostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) WithHTTPClient(client *http.Client) *PostPublicationsPublicationIDPostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) WithBody(body *models.Post) *PostPublicationsPublicationIDPostsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) SetBody(body *models.Post) {
	o.Body = body
}

// WithPublicationID adds the publicationID to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) WithPublicationID(publicationID string) *PostPublicationsPublicationIDPostsParams {
	o.SetPublicationID(publicationID)
	return o
}

// SetPublicationID adds the publicationId to the post publications publication ID posts params
func (o *PostPublicationsPublicationIDPostsParams) SetPublicationID(publicationID string) {
	o.PublicationID = publicationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicationsPublicationIDPostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param publicationId
	if err := r.SetPathParam("publicationId", o.PublicationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}