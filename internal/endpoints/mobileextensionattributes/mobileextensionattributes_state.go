// mobileextensionattributes_state.go
package mobileextensionattributes

import (
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// updateTerraformState updates the Terraform state with the latest Mobile Extension Attribute information from the Jamf Pro API.
func updateTerraformState(d *schema.ResourceData, resource *jamfpro.ResourceMobileExtensionAttribute) diag.Diagnostics {
	var diags diag.Diagnostics

	// Update the Terraform state with the fetched data
	resourceData := map[string]interface{}{
		"name":              resource.Name,
		"description":       resource.Description,
		"data_type":         strings.ToLower(resource.DataType),
		"inventory_display": resource.InventoryDisplay,
		"input_type": []interface{}{
			map[string]interface{}{
				"type":     resource.InputType.Type,
			},
		},
	}

	for key, val := range resourceData {
		if err := d.Set(key, val); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags

}
