package certificate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CertificateId{}

// CertificateId is a struct representing the Resource ID for a Certificate
type CertificateId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	CertificateId     string
}

// NewCertificateID returns a new CertificateId struct
func NewCertificateID(subscriptionId string, resourceGroupName string, serviceName string, certificateId string) CertificateId {
	return CertificateId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		CertificateId:     certificateId,
	}
}

// ParseCertificateID parses 'input' into a CertificateId
func ParseCertificateID(input string) (*CertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(CertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.CertificateId, ok = parsed.Parsed["certificateId"]; !ok {
		return nil, fmt.Errorf("the segment 'certificateId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCertificateIDInsensitively parses 'input' case-insensitively into a CertificateId
// note: this method should only be used for API response data and not user input
func ParseCertificateIDInsensitively(input string) (*CertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(CertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.CertificateId, ok = parsed.Parsed["certificateId"]; !ok {
		return nil, fmt.Errorf("the segment 'certificateId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCertificateID checks that 'input' can be parsed as a Certificate ID
func ValidateCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Certificate ID
func (id CertificateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/certificates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.CertificateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Certificate ID
func (id CertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticCertificates", "certificates", "certificates"),
		resourceids.UserSpecifiedSegment("certificateId", "certificateIdValue"),
	}
}

// String returns a human-readable description of this Certificate ID
func (id CertificateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Certificate: %q", id.CertificateId),
	}
	return fmt.Sprintf("Certificate (%s)", strings.Join(components, "\n"))
}
