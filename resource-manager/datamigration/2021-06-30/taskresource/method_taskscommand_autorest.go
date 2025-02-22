package taskresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksCommandOperationResponse struct {
	HttpResponse *http.Response
	Model        *CommandProperties
}

// TasksCommand ...
func (c TaskResourceClient) TasksCommand(ctx context.Context, id TaskId, input CommandProperties) (result TasksCommandOperationResponse, err error) {
	req, err := c.preparerForTasksCommand(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksCommand", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksCommand", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTasksCommand(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksCommand", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTasksCommand prepares the TasksCommand request.
func (c TaskResourceClient) preparerForTasksCommand(ctx context.Context, id TaskId, input CommandProperties) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/command", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTasksCommand handles the response to the TasksCommand request. The method always
// closes the http.Response Body.
func (c TaskResourceClient) responderForTasksCommand(resp *http.Response) (result TasksCommandOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
