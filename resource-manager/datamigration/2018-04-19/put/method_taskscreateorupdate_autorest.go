package put

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectTask
}

// TasksCreateOrUpdate ...
func (c PUTClient) TasksCreateOrUpdate(ctx context.Context, id TaskId, input ProjectTask) (result TasksCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForTasksCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "TasksCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "TasksCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTasksCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "TasksCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTasksCreateOrUpdate prepares the TasksCreateOrUpdate request.
func (c PUTClient) preparerForTasksCreateOrUpdate(ctx context.Context, id TaskId, input ProjectTask) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTasksCreateOrUpdate handles the response to the TasksCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c PUTClient) responderForTasksCreateOrUpdate(resp *http.Response) (result TasksCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
