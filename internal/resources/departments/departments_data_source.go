// department_data_source.go
package departments

import (
	"context"
	"fmt"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceJamfProDepartments provides information about a specific department in Jamf Pro.
func DataSourceJamfProDepartments() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceJamfProDepartmentsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The unique identifier of the department.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique name of the jamf pro department.",
				Computed:    true,
			},
		},
	}
}

// DataSourceJamfProDepartmentsRead fetches the details of a specific department from Jamf Pro using either its unique Name or its Id.
func DataSourceJamfProDepartmentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*client.APIClient).Conn

	var department *jamfpro.ResponseDepartment
	var err error

	// Check if Name is provided in the data source configuration
	if v, ok := d.GetOk("name"); ok && v.(string) != "" {
		departmentName := v.(string)
		department, err = conn.GetDepartmentByName(departmentName)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to fetch department by name: %v", err))
		}
	} else if v, ok := d.GetOk("id"); ok {
		departmentID, convertErr := strconv.Atoi(v.(string))
		if convertErr != nil {
			return diag.FromErr(fmt.Errorf("failed to convert department ID to integer: %v", convertErr))
		}
		department, err = conn.GetDepartmentByID(departmentID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to fetch department by ID: %v", err))
		}
	} else {
		return diag.Errorf("Either 'name' or 'id' must be provided")
	}

	if department == nil {
		return diag.FromErr(fmt.Errorf("department not found"))
	}

	// Set the data source attributes using the fetched data
	d.SetId(fmt.Sprintf("%d", department.ID))
	d.Set("name", department.Name)

	return nil
}