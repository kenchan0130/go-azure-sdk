package vmhost

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

type SubAccountListVMHostsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMResources

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SubAccountListVMHostsOperationResponse, error)
}

type SubAccountListVMHostsCompleteResult struct {
	Items []VMResources
}

func (r SubAccountListVMHostsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SubAccountListVMHostsOperationResponse) LoadMore(ctx context.Context) (resp SubAccountListVMHostsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SubAccountListVMHosts ...
func (c VMHostClient) SubAccountListVMHosts(ctx context.Context, id AccountId) (resp SubAccountListVMHostsOperationResponse, err error) {
	req, err := c.preparerForSubAccountListVMHosts(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHosts", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHosts", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSubAccountListVMHosts(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHosts", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSubAccountListVMHosts prepares the SubAccountListVMHosts request.
func (c VMHostClient) preparerForSubAccountListVMHosts(ctx context.Context, id AccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listVMHosts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSubAccountListVMHostsWithNextLink prepares the SubAccountListVMHosts request with the given nextLink token.
func (c VMHostClient) preparerForSubAccountListVMHostsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSubAccountListVMHosts handles the response to the SubAccountListVMHosts request. The method always
// closes the http.Response Body.
func (c VMHostClient) responderForSubAccountListVMHosts(resp *http.Response) (result SubAccountListVMHostsOperationResponse, err error) {
	type page struct {
		Values   []VMResources `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SubAccountListVMHostsOperationResponse, err error) {
			req, err := c.preparerForSubAccountListVMHostsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHosts", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHosts", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSubAccountListVMHosts(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHosts", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SubAccountListVMHostsComplete retrieves all of the results into a single object
func (c VMHostClient) SubAccountListVMHostsComplete(ctx context.Context, id AccountId) (SubAccountListVMHostsCompleteResult, error) {
	return c.SubAccountListVMHostsCompleteMatchingPredicate(ctx, id, VMResourcesOperationPredicate{})
}

// SubAccountListVMHostsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VMHostClient) SubAccountListVMHostsCompleteMatchingPredicate(ctx context.Context, id AccountId, predicate VMResourcesOperationPredicate) (resp SubAccountListVMHostsCompleteResult, err error) {
	items := make([]VMResources, 0)

	page, err := c.SubAccountListVMHosts(ctx, id)
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

	out := SubAccountListVMHostsCompleteResult{
		Items: items,
	}
	return out, nil
}
