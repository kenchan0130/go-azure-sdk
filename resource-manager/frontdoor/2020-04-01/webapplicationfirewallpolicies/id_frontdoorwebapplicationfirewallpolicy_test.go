package webapplicationfirewallpolicies

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FrontDoorWebApplicationFirewallPolicyId{}

func TestNewFrontDoorWebApplicationFirewallPolicyID(t *testing.T) {
	id := NewFrontDoorWebApplicationFirewallPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorWebApplicationFirewallPolicyValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.FrontDoorWebApplicationFirewallPolicyName != "frontDoorWebApplicationFirewallPolicyValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FrontDoorWebApplicationFirewallPolicyName'", id.FrontDoorWebApplicationFirewallPolicyName, "frontDoorWebApplicationFirewallPolicyValue")
	}
}

func TestFormatFrontDoorWebApplicationFirewallPolicyID(t *testing.T) {
	actual := NewFrontDoorWebApplicationFirewallPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorWebApplicationFirewallPolicyValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/frontDoorWebApplicationFirewallPolicyValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseFrontDoorWebApplicationFirewallPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FrontDoorWebApplicationFirewallPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/frontDoorWebApplicationFirewallPolicyValue",
			Expected: &FrontDoorWebApplicationFirewallPolicyId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				FrontDoorWebApplicationFirewallPolicyName: "frontDoorWebApplicationFirewallPolicyValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/frontDoorWebApplicationFirewallPolicyValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFrontDoorWebApplicationFirewallPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.FrontDoorWebApplicationFirewallPolicyName != v.Expected.FrontDoorWebApplicationFirewallPolicyName {
			t.Fatalf("Expected %q but got %q for FrontDoorWebApplicationFirewallPolicyName", v.Expected.FrontDoorWebApplicationFirewallPolicyName, actual.FrontDoorWebApplicationFirewallPolicyName)
		}

	}
}

func TestParseFrontDoorWebApplicationFirewallPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FrontDoorWebApplicationFirewallPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/fRoNtDoOrWeBaPpLiCaTiOnFiReWaLlPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/frontDoorWebApplicationFirewallPolicyValue",
			Expected: &FrontDoorWebApplicationFirewallPolicyId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				FrontDoorWebApplicationFirewallPolicyName: "frontDoorWebApplicationFirewallPolicyValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/frontDoorWebApplicationFirewallPolicyValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/fRoNtDoOrWeBaPpLiCaTiOnFiReWaLlPoLiCiEs/fRoNtDoOrWeBaPpLiCaTiOnFiReWaLlPoLiCyVaLuE",
			Expected: &FrontDoorWebApplicationFirewallPolicyId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "eXaMpLe-rEsOuRcE-GrOuP",
				FrontDoorWebApplicationFirewallPolicyName: "fRoNtDoOrWeBaPpLiCaTiOnFiReWaLlPoLiCyVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/fRoNtDoOrWeBaPpLiCaTiOnFiReWaLlPoLiCiEs/fRoNtDoOrWeBaPpLiCaTiOnFiReWaLlPoLiCyVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFrontDoorWebApplicationFirewallPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.FrontDoorWebApplicationFirewallPolicyName != v.Expected.FrontDoorWebApplicationFirewallPolicyName {
			t.Fatalf("Expected %q but got %q for FrontDoorWebApplicationFirewallPolicyName", v.Expected.FrontDoorWebApplicationFirewallPolicyName, actual.FrontDoorWebApplicationFirewallPolicyName)
		}

	}
}

func TestSegmentsForFrontDoorWebApplicationFirewallPolicyId(t *testing.T) {
	segments := FrontDoorWebApplicationFirewallPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("FrontDoorWebApplicationFirewallPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
