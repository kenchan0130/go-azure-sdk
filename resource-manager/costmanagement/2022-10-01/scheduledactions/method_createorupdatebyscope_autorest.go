package scheduledactions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateByScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *ScheduledAction
}

type CreateOrUpdateByScopeOperationOptions struct {
	IfMatch *string
}

func DefaultCreateOrUpdateByScopeOperationOptions() CreateOrUpdateByScopeOperationOptions {
	return CreateOrUpdateByScopeOperationOptions{}
}

func (o CreateOrUpdateByScopeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o CreateOrUpdateByScopeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// CreateOrUpdateByScope ...
func (c ScheduledActionsClient) CreateOrUpdateByScope(ctx context.Context, id ScopedScheduledActionId, input ScheduledAction, options CreateOrUpdateByScopeOperationOptions) (result CreateOrUpdateByScopeOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateByScope(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "CreateOrUpdateByScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "CreateOrUpdateByScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateByScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "CreateOrUpdateByScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateByScope prepares the CreateOrUpdateByScope request.
func (c ScheduledActionsClient) preparerForCreateOrUpdateByScope(ctx context.Context, id ScopedScheduledActionId, input ScheduledAction, options CreateOrUpdateByScopeOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateByScope handles the response to the CreateOrUpdateByScope request. The method always
// closes the http.Response Body.
func (c ScheduledActionsClient) responderForCreateOrUpdateByScope(resp *http.Response) (result CreateOrUpdateByScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
