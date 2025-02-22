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

type GetConnectionInfoOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeConnectionInfo
}

// GetConnectionInfo ...
func (c IntegrationRuntimesClient) GetConnectionInfo(ctx context.Context, id IntegrationRuntimeId) (result GetConnectionInfoOperationResponse, err error) {
	req, err := c.preparerForGetConnectionInfo(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "GetConnectionInfo", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "GetConnectionInfo", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetConnectionInfo(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "GetConnectionInfo", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetConnectionInfo prepares the GetConnectionInfo request.
func (c IntegrationRuntimesClient) preparerForGetConnectionInfo(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getConnectionInfo", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetConnectionInfo handles the response to the GetConnectionInfo request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForGetConnectionInfo(resp *http.Response) (result GetConnectionInfoOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
