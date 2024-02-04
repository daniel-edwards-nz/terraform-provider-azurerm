// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type ManagementGroupCostManagementViewId struct {
	ManagementGroupName string
	ViewName            string
}

func NewManagementGroupCostManagementViewID(managementGroupName, viewName string) ManagementGroupCostManagementViewId {
	return ManagementGroupCostManagementViewId{
		ManagementGroupName: managementGroupName,
		ViewName:            viewName,
	}
}

func (id ManagementGroupCostManagementViewId) String() string {
	segments := []string{
		fmt.Sprintf("View Name %q", id.ViewName),
		fmt.Sprintf("Management Group Name %q", id.ManagementGroupName),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Management Group Cost Management View", segmentsStr)
}

func (id ManagementGroupCostManagementViewId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.CostManagement/views/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupName, id.ViewName)
}

// ManagementGroupCostManagementViewID parses a ManagementGroupCostManagementView ID into an ManagementGroupCostManagementViewId struct
func ManagementGroupCostManagementViewID(input string) (*ManagementGroupCostManagementViewId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("parsing %q as an ManagementGroupCostManagementView ID: %+v", input, err)
	}

	resourceId := ManagementGroupCostManagementViewId{}

	if resourceId.ManagementGroupName, err = id.PopSegment("managementGroups"); err != nil {
		return nil, err
	}
	if resourceId.ViewName, err = id.PopSegment("views"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
