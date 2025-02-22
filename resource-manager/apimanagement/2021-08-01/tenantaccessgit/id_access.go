package tenantaccessgit

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AccessId{}

// AccessId is a struct representing the Resource ID for a Access
type AccessId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	AccessName        AccessIdName
}

// NewAccessID returns a new AccessId struct
func NewAccessID(subscriptionId string, resourceGroupName string, serviceName string, accessName AccessIdName) AccessId {
	return AccessId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		AccessName:        accessName,
	}
}

// ParseAccessID parses 'input' into a AccessId
func ParseAccessID(input string) (*AccessId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccessId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccessId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["accessName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'accessName' was not found in the resource id %q", input)
		}

		accessName, err := parseAccessIdName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.AccessName = *accessName
	}

	return &id, nil
}

// ParseAccessIDInsensitively parses 'input' case-insensitively into a AccessId
// note: this method should only be used for API response data and not user input
func ParseAccessIDInsensitively(input string) (*AccessId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccessId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccessId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["accessName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'accessName' was not found in the resource id %q", input)
		}

		accessName, err := parseAccessIdName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.AccessName = *accessName
	}

	return &id, nil
}

// ValidateAccessID checks that 'input' can be parsed as a Access ID
func ValidateAccessID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAccessID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Access ID
func (id AccessId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/tenant/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, string(id.AccessName))
}

// Segments returns a slice of Resource ID Segments which comprise this Access ID
func (id AccessId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticTenant", "tenant", "tenant"),
		resourceids.ConstantSegment("accessName", PossibleValuesForAccessIdName(), "access"),
	}
}

// String returns a human-readable description of this Access ID
func (id AccessId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Access Name: %q", string(id.AccessName)),
	}
	return fmt.Sprintf("Access (%s)", strings.Join(components, "\n"))
}
