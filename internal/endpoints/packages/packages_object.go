// packages_data_object.go
package packages

import (
	"encoding/xml"
	"fmt"
	"log"
	"path/filepath"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// constructJamfProPackageCreate constructs a ResourcePackage object from the provided schema data.
// It extracts the filename from the full path provided in the schema and uses it for the Filename field.
func constructJamfProPackageCreate(d *schema.ResourceData) (*jamfpro.ResourcePackage, error) {
	// Extract the full file path from the schema
	fullPath := d.Get("package_file_path").(string)

	// Use filepath.Base to extract just the filename from the full path
	filename := filepath.Base(fullPath)

	// Get the category from the schema, and set it to "Unknown" if it's empty
	// 'Unknown' is the valid default request value for the category field when none is set
	// Jamf API returns "No category assigned" for the same field. But this is not a valid
	// request value. Why!!!! >_<
	category := d.Get("category").(string)
	if category == "" {
		category = "Unknown"
	}

	packageResource := &jamfpro.ResourcePackage{
		Name:               d.Get("name").(string),
		Filename:           filename,
		Category:           category,
		Info:               d.Get("info").(string),
		Notes:              d.Get("notes").(string),
		Priority:           d.Get("priority").(int),
		RebootRequired:     d.Get("reboot_required").(bool),
		FillUserTemplate:   d.Get("fill_user_template").(bool),
		FillExistingUsers:  d.Get("fill_existing_users").(bool),
		BootVolumeRequired: d.Get("boot_volume_required").(bool),
		AllowUninstalled:   d.Get("allow_uninstalled").(bool),
		OSRequirements:     d.Get("os_requirements").(string),
		// fields appear to only be relevant for jamf admin indexed packages
		// which i believe is to be deprecated.
		//RequiredProcessor:          d.Get("required_processor").(string),
		//SwitchWithPackage:          d.Get("switch_with_package").(string),
		//ReinstallOption:            d.Get("reinstall_option").(string),
		//TriggeringFiles:            d.Get("triggering_files").(string),
		InstallIfReportedAvailable: d.Get("install_if_reported_available").(bool),
		SendNotification:           d.Get("send_notification").(bool),
	}

	// Serialize and pretty-print the Package object as XML for logging
	resourceXML, err := xml.MarshalIndent(packageResource, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Jamf Pro Package '%s' to XML: %v", packageResource.Name, err)
	}

	log.Printf("[DEBUG] Constructed Jamf Pro Package XML:\n%s\n", string(resourceXML))

	return packageResource, nil
}
