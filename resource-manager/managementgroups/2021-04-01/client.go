package v2021_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/entities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/managementgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/tenantbackfill"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	CheckNameAvailability *checknameavailability.CheckNameAvailabilityClient
	Entities              *entities.EntitiesClient
	ManagementGroups      *managementgroups.ManagementGroupsClient
	TenantBackfill        *tenantbackfill.TenantBackfillClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	checkNameAvailabilityClient := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkNameAvailabilityClient.Client)

	entitiesClient := entities.NewEntitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&entitiesClient.Client)

	managementGroupsClient := managementgroups.NewManagementGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&managementGroupsClient.Client)

	tenantBackfillClient := tenantbackfill.NewTenantBackfillClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantBackfillClient.Client)

	return Client{
		CheckNameAvailability: &checkNameAvailabilityClient,
		Entities:              &entitiesClient,
		ManagementGroups:      &managementGroupsClient,
		TenantBackfill:        &tenantBackfillClient,
	}
}
