package v2020_10_05_privatepreview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servers"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Configurations *configurations.ConfigurationsClient
	FirewallRules  *firewallrules.FirewallRulesClient
	Roles          *roles.RolesClient
	ServerGroups   *servergroups.ServerGroupsClient
	Servers        *servers.ServersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	configurationsClient := configurations.NewConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	rolesClient := roles.NewRolesClientWithBaseURI(endpoint)
	configureAuthFunc(&rolesClient.Client)

	serverGroupsClient := servergroups.NewServerGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&serverGroupsClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	return Client{
		Configurations: &configurationsClient,
		FirewallRules:  &firewallRulesClient,
		Roles:          &rolesClient,
		ServerGroups:   &serverGroupsClient,
		Servers:        &serversClient,
	}
}
