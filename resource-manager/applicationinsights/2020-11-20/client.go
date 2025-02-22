package v2020_11_20

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2020-11-20/workbooktemplatesapis"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	WorkbookTemplatesAPIs *workbooktemplatesapis.WorkbookTemplatesAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	workbookTemplatesAPIsClient := workbooktemplatesapis.NewWorkbookTemplatesAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&workbookTemplatesAPIsClient.Client)

	return Client{
		WorkbookTemplatesAPIs: &workbookTemplatesAPIsClient,
	}
}
