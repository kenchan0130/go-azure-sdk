package v2022_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-06-01/policyassignments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	PolicyAssignments *policyassignments.PolicyAssignmentsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	policyAssignmentsClient := policyassignments.NewPolicyAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&policyAssignmentsClient.Client)

	return Client{
		PolicyAssignments: &policyAssignmentsClient,
	}
}
