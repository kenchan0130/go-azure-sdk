package privateendpointconnection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type UpdateOperationOptions struct {
	IfMatch *string
}

func DefaultUpdateOperationOptions() UpdateOperationOptions {
	return UpdateOperationOptions{}
}

func (o UpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o UpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// Update ...
func (c PrivateEndpointConnectionClient) Update(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection, options UpdateOperationOptions) (result UpdateOperationResponse, err error) {
	req, err := c.preparerForUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "Update", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateThenPoll performs Update then polls until it's completed
func (c PrivateEndpointConnectionClient) UpdateThenPoll(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection, options UpdateOperationOptions) error {
	result, err := c.Update(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing Update: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Update: %+v", err)
	}

	return nil
}

// preparerForUpdate prepares the Update request.
func (c PrivateEndpointConnectionClient) preparerForUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection, options UpdateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdate sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (c PrivateEndpointConnectionClient) senderForUpdate(ctx context.Context, req *http.Request) (future UpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
