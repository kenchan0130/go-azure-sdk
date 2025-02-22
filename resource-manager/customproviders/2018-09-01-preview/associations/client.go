package associations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssociationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAssociationsClientWithBaseURI(endpoint string) AssociationsClient {
	return AssociationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
