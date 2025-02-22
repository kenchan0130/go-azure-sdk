package managedprivateendpoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ManagedPrivateEndpointId{}

// ManagedPrivateEndpointId is a struct representing the Resource ID for a Managed Private Endpoint
type ManagedPrivateEndpointId struct {
	SubscriptionId             string
	ResourceGroupName          string
	FactoryName                string
	ManagedVirtualNetworkName  string
	ManagedPrivateEndpointName string
}

// NewManagedPrivateEndpointID returns a new ManagedPrivateEndpointId struct
func NewManagedPrivateEndpointID(subscriptionId string, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string) ManagedPrivateEndpointId {
	return ManagedPrivateEndpointId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		FactoryName:                factoryName,
		ManagedVirtualNetworkName:  managedVirtualNetworkName,
		ManagedPrivateEndpointName: managedPrivateEndpointName,
	}
}

// ParseManagedPrivateEndpointID parses 'input' into a ManagedPrivateEndpointId
func ParseManagedPrivateEndpointID(input string) (*ManagedPrivateEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedPrivateEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedPrivateEndpointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.ManagedVirtualNetworkName, ok = parsed.Parsed["managedVirtualNetworkName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedVirtualNetworkName' was not found in the resource id %q", input)
	}

	if id.ManagedPrivateEndpointName, ok = parsed.Parsed["managedPrivateEndpointName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedPrivateEndpointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseManagedPrivateEndpointIDInsensitively parses 'input' case-insensitively into a ManagedPrivateEndpointId
// note: this method should only be used for API response data and not user input
func ParseManagedPrivateEndpointIDInsensitively(input string) (*ManagedPrivateEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedPrivateEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedPrivateEndpointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.ManagedVirtualNetworkName, ok = parsed.Parsed["managedVirtualNetworkName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedVirtualNetworkName' was not found in the resource id %q", input)
	}

	if id.ManagedPrivateEndpointName, ok = parsed.Parsed["managedPrivateEndpointName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedPrivateEndpointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateManagedPrivateEndpointID checks that 'input' can be parsed as a Managed Private Endpoint ID
func ValidateManagedPrivateEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedPrivateEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Private Endpoint ID
func (id ManagedPrivateEndpointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/managedVirtualNetworks/%s/managedPrivateEndpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.ManagedVirtualNetworkName, id.ManagedPrivateEndpointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Private Endpoint ID
func (id ManagedPrivateEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticManagedVirtualNetworks", "managedVirtualNetworks", "managedVirtualNetworks"),
		resourceids.UserSpecifiedSegment("managedVirtualNetworkName", "managedVirtualNetworkValue"),
		resourceids.StaticSegment("staticManagedPrivateEndpoints", "managedPrivateEndpoints", "managedPrivateEndpoints"),
		resourceids.UserSpecifiedSegment("managedPrivateEndpointName", "managedPrivateEndpointValue"),
	}
}

// String returns a human-readable description of this Managed Private Endpoint ID
func (id ManagedPrivateEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Managed Virtual Network Name: %q", id.ManagedVirtualNetworkName),
		fmt.Sprintf("Managed Private Endpoint Name: %q", id.ManagedPrivateEndpointName),
	}
	return fmt.Sprintf("Managed Private Endpoint (%s)", strings.Join(components, "\n"))
}
