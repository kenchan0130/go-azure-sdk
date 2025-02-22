package authorizationloginlinks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostOperationResponse struct {
	HttpResponse *http.Response
	Model        *AuthorizationLoginResponseContract
}

// Post ...
func (c AuthorizationLoginLinksClient) Post(ctx context.Context, id AuthorizationId, input AuthorizationLoginRequestContract) (result PostOperationResponse, err error) {
	req, err := c.preparerForPost(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationloginlinks.AuthorizationLoginLinksClient", "Post", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationloginlinks.AuthorizationLoginLinksClient", "Post", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPost(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationloginlinks.AuthorizationLoginLinksClient", "Post", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPost prepares the Post request.
func (c AuthorizationLoginLinksClient) preparerForPost(ctx context.Context, id AuthorizationId, input AuthorizationLoginRequestContract) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getLoginLinks", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPost handles the response to the Post request. The method always
// closes the http.Response Body.
func (c AuthorizationLoginLinksClient) responderForPost(resp *http.Response) (result PostOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
