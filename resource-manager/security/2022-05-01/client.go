package v2022_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2022-05-01/settings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Settings *settings.SettingsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	settingsClient := settings.NewSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&settingsClient.Client)

	return Client{
		Settings: &settingsClient,
	}
}
