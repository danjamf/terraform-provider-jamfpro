// allowedfileextensions_object.go
package allowedfileextensions

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// constructJamfProAllowedFileExtension creates a new ResourceAllowedFileExtension instance from Terraform data and serializes it to XML.
func constructJamfProAllowedFileExtension(d *schema.ResourceData) (*jamfpro.ResourceAllowedFileExtension, error) {
	allowedFileExtension := &jamfpro.ResourceAllowedFileExtension{
		Extension: d.Get("extension").(string),
	}

	// Serialize and pretty-print the Allowed File Extension object as XML for logging
	resourceXML, err := xml.MarshalIndent(allowedFileExtension, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Jamf Pro Allowed File Extension '%s' to XML: %v", allowedFileExtension.Extension, err)
	}

	// Use log.Printf instead of fmt.Printf for logging within the Terraform provider context
	log.Printf("[DEBUG] Constructed Jamf Pro Allowed File Extension XML:\n%s\n", string(resourceXML))

	return allowedFileExtension, nil
}
