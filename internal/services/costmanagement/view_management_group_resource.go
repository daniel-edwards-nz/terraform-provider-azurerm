// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package costmanagement

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/costmanagement/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type ManagementGroupCostManagementViewResource struct {
	base costManagementViewBaseResource
}

var _ sdk.Resource = ManagementGroupCostManagementViewResource{}

func (r ManagementGroupCostManagementViewResource) Arguments() map[string]*pluginsdk.Schema {
	schema := map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotWhiteSpace,
		},
		"management_group_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: commonids.ValidateManagementGroupID,
		},
	}
	return r.base.arguments(schema)
}

func (r ManagementGroupCostManagementViewResource) Attributes() map[string]*pluginsdk.Schema {
	return r.base.attributes()
}

func (r ManagementGroupCostManagementViewResource) ModelObject() interface{} {
	return nil
}

func (r ManagementGroupCostManagementViewResource) ResourceType() string {
	return "azurerm_management_group_cost_management_view"
}

func (r ManagementGroupCostManagementViewResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validate.ManagementGroupCostManagementViewID
}

func (r ManagementGroupCostManagementViewResource) Create() sdk.ResourceFunc {
	return r.base.createFunc(r.ResourceType(), "management_group_id")
}

func (r ManagementGroupCostManagementViewResource) Read() sdk.ResourceFunc {
	return r.base.readFunc("management_group_id")
}

func (r ManagementGroupCostManagementViewResource) Delete() sdk.ResourceFunc {
	return r.base.deleteFunc()
}

func (r ManagementGroupCostManagementViewResource) Update() sdk.ResourceFunc {
	return r.base.updateFunc()
}
