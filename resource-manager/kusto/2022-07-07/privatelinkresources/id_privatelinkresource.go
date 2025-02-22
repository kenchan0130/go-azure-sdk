package privatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PrivateLinkResourceId{}

// PrivateLinkResourceId is a struct representing the Resource ID for a Private Link Resource
type PrivateLinkResourceId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ClusterName             string
	PrivateLinkResourceName string
}

// NewPrivateLinkResourceID returns a new PrivateLinkResourceId struct
func NewPrivateLinkResourceID(subscriptionId string, resourceGroupName string, clusterName string, privateLinkResourceName string) PrivateLinkResourceId {
	return PrivateLinkResourceId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ClusterName:             clusterName,
		PrivateLinkResourceName: privateLinkResourceName,
	}
}

// ParsePrivateLinkResourceID parses 'input' into a PrivateLinkResourceId
func ParsePrivateLinkResourceID(input string) (*PrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.PrivateLinkResourceName, ok = parsed.Parsed["privateLinkResourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkResourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a PrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParsePrivateLinkResourceIDInsensitively(input string) (*PrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.PrivateLinkResourceName, ok = parsed.Parsed["privateLinkResourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkResourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePrivateLinkResourceID checks that 'input' can be parsed as a Private Link Resource ID
func ValidatePrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Private Link Resource ID
func (id PrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Kusto/clusters/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.PrivateLinkResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Private Link Resource ID
func (id PrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKusto", "Microsoft.Kusto", "Microsoft.Kusto"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("privateLinkResourceName", "privateLinkResourceValue"),
	}
}

// String returns a human-readable description of this Private Link Resource ID
func (id PrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Private Link Resource Name: %q", id.PrivateLinkResourceName),
	}
	return fmt.Sprintf("Private Link Resource (%s)", strings.Join(components, "\n"))
}
