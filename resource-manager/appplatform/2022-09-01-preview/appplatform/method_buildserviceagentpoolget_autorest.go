package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildServiceAgentPoolGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *BuildServiceAgentPoolResource
}

// BuildServiceAgentPoolGet ...
func (c AppPlatformClient) BuildServiceAgentPoolGet(ctx context.Context, id AgentPoolId) (result BuildServiceAgentPoolGetOperationResponse, err error) {
	req, err := c.preparerForBuildServiceAgentPoolGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceAgentPoolGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceAgentPoolGet prepares the BuildServiceAgentPoolGet request.
func (c AppPlatformClient) preparerForBuildServiceAgentPoolGet(ctx context.Context, id AgentPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBuildServiceAgentPoolGet handles the response to the BuildServiceAgentPoolGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceAgentPoolGet(resp *http.Response) (result BuildServiceAgentPoolGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
