package appplatform

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StoragesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StorageResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (StoragesListOperationResponse, error)
}

type StoragesListCompleteResult struct {
	Items []StorageResource
}

func (r StoragesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r StoragesListOperationResponse) LoadMore(ctx context.Context) (resp StoragesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// StoragesList ...
func (c AppPlatformClient) StoragesList(ctx context.Context, id SpringId) (resp StoragesListOperationResponse, err error) {
	req, err := c.preparerForStoragesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForStoragesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForStoragesList prepares the StoragesList request.
func (c AppPlatformClient) preparerForStoragesList(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/storages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForStoragesListWithNextLink prepares the StoragesList request with the given nextLink token.
func (c AppPlatformClient) preparerForStoragesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStoragesList handles the response to the StoragesList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForStoragesList(resp *http.Response) (result StoragesListOperationResponse, err error) {
	type page struct {
		Values   []StorageResource `json:"value"`
		NextLink *string           `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result StoragesListOperationResponse, err error) {
			req, err := c.preparerForStoragesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForStoragesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// StoragesListComplete retrieves all of the results into a single object
func (c AppPlatformClient) StoragesListComplete(ctx context.Context, id SpringId) (StoragesListCompleteResult, error) {
	return c.StoragesListCompleteMatchingPredicate(ctx, id, StorageResourceOperationPredicate{})
}

// StoragesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) StoragesListCompleteMatchingPredicate(ctx context.Context, id SpringId, predicate StorageResourceOperationPredicate) (resp StoragesListCompleteResult, err error) {
	items := make([]StorageResource, 0)

	page, err := c.StoragesList(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := StoragesListCompleteResult{
		Items: items,
	}
	return out, nil
}
