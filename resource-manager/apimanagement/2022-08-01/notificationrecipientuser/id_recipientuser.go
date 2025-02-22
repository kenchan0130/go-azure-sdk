package notificationrecipientuser

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RecipientUserId{}

// RecipientUserId is a struct representing the Resource ID for a Recipient User
type RecipientUserId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	NotificationName  NotificationName
	UserId            string
}

// NewRecipientUserID returns a new RecipientUserId struct
func NewRecipientUserID(subscriptionId string, resourceGroupName string, serviceName string, notificationName NotificationName, userId string) RecipientUserId {
	return RecipientUserId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		NotificationName:  notificationName,
		UserId:            userId,
	}
}

// ParseRecipientUserID parses 'input' into a RecipientUserId
func ParseRecipientUserID(input string) (*RecipientUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecipientUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecipientUserId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["notificationName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'notificationName' was not found in the resource id %q", input)
		}

		notificationName, err := parseNotificationName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.NotificationName = *notificationName
	}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, fmt.Errorf("the segment 'userId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRecipientUserIDInsensitively parses 'input' case-insensitively into a RecipientUserId
// note: this method should only be used for API response data and not user input
func ParseRecipientUserIDInsensitively(input string) (*RecipientUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecipientUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecipientUserId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["notificationName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'notificationName' was not found in the resource id %q", input)
		}

		notificationName, err := parseNotificationName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.NotificationName = *notificationName
	}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, fmt.Errorf("the segment 'userId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRecipientUserID checks that 'input' can be parsed as a Recipient User ID
func ValidateRecipientUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecipientUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recipient User ID
func (id RecipientUserId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/notifications/%s/recipientUsers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, string(id.NotificationName), id.UserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Recipient User ID
func (id RecipientUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticNotifications", "notifications", "notifications"),
		resourceids.ConstantSegment("notificationName", PossibleValuesForNotificationName(), "AccountClosedPublisher"),
		resourceids.StaticSegment("staticRecipientUsers", "recipientUsers", "recipientUsers"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
	}
}

// String returns a human-readable description of this Recipient User ID
func (id RecipientUserId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Notification Name: %q", string(id.NotificationName)),
		fmt.Sprintf("User: %q", id.UserId),
	}
	return fmt.Sprintf("Recipient User (%s)", strings.Join(components, "\n"))
}
