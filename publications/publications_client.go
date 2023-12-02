// Code generated by go-swagger; DO NOT EDIT.

package publications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new publications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for publications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPublicationsPublicationIDContributors(params *GetPublicationsPublicationIDContributorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublicationsPublicationIDContributorsOK, error)

	GetUsersUserIDPublications(params *GetUsersUserIDPublicationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersUserIDPublicationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetPublicationsPublicationIDContributors contributors of publication

This endpoint returns a list of contributors for a given publication. In other words, a list of Medium users who are allowed to publish under a publication, as well as a description of their exact role in the publication (for now, either an editor or a writer).
*/
func (a *Client) GetPublicationsPublicationIDContributors(params *GetPublicationsPublicationIDContributorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublicationsPublicationIDContributorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicationsPublicationIDContributorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPublicationsPublicationIDContributors",
		Method:             "GET",
		PathPattern:        "/publications/{publicationId}/contributors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublicationsPublicationIDContributorsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPublicationsPublicationIDContributorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPublicationsPublicationIDContributors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUsersUserIDPublications users s publications

Returns a full list of publications that the user is related to in some way. This includes all publications the user is subscribed to, writes to, or edits.
*/
func (a *Client) GetUsersUserIDPublications(params *GetUsersUserIDPublicationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersUserIDPublicationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDPublicationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUsersUserIDPublications",
		Method:             "GET",
		PathPattern:        "/users/{userId}/publications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDPublicationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDPublicationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUserIDPublications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
