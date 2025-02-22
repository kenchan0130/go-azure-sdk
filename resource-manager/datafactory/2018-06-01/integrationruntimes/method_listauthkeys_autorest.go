package integrationruntimes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAuthKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeAuthKeys
}

// ListAuthKeys ...
func (c IntegrationRuntimesClient) ListAuthKeys(ctx context.Context, id IntegrationRuntimeId) (result ListAuthKeysOperationResponse, err error) {
	req, err := c.preparerForListAuthKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "ListAuthKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "ListAuthKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListAuthKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "ListAuthKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListAuthKeys prepares the ListAuthKeys request.
func (c IntegrationRuntimesClient) preparerForListAuthKeys(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listAuthKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListAuthKeys handles the response to the ListAuthKeys request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForListAuthKeys(resp *http.Response) (result ListAuthKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
