// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = ManagementGroupCostManagementViewId{}

func TestManagementGroupCostManagementViewIDFormatter(t *testing.T) {
	actual := NewManagementGroupCostManagementViewID("12345678-1234-9876-4563-123456789012", "view1").ID()
	expected := "/providers/Microsoft.Management/managementGroups/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/views/view1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestManagementGroupCostManagementViewID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagementGroupCostManagementViewId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing ManagementGroupName
			Input: "/providers/Microsoft.Management/",
			Error: true,
		},

		{
			// missing value for ManagementGroupName
			Input: "/providers/Microsoft.Management/managementGroups/",
			Error: true,
		},

		{
			// missing ViewName
			Input: "/providers/Microsoft.Management/managementGroups/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/",
			Error: true,
		},

		{
			// missing value for ViewName
			Input: "/providers/Microsoft.Management/managementGroups/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/views/",
			Error: true,
		},

		{
			// valid
			Input: "/providers/Microsoft.Management/managementGroups/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/views/view1",
			Expected: &ManagementGroupCostManagementViewId{
				ManagementGroupName: "12345678-1234-9876-4563-123456789012",
				ViewName:            "view1",
			},
		},

		{
			// upper-cased
			Input: "/PROVIDERS/MICROSOFT.MANAGEMENT/MANAGEMENTGROUPS/12345678-1234-9876-4563-123456789012/PROVIDERS/MICROSOFT.COSTMANAGEMENT/VIEWS/VIEW1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ManagementGroupCostManagementViewID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupName != v.Expected.ManagementGroupName {
			t.Fatalf("Expected %q but got %q for ManagementGroupName", v.Expected.ManagementGroupName, actual.ManagementGroupName)
		}
		if actual.ViewName != v.Expected.ViewName {
			t.Fatalf("Expected %q but got %q for ViewName", v.Expected.ViewName, actual.ViewName)
		}
	}
}
