// mobileextensionattributes_object.go
package mobileextensionattributes

import (
	"encoding/xml"
	"fmt"
	"log"
	//"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// constructJamfProMobileExtensionAttribute constructs a ResourceMobileExtensionAttribute object from the provided schema data.
func constructJamfProMobileExtensionAttribute(d *schema.ResourceData) (*jamfpro.ResourceMobileExtensionAttribute, error) {
	attribute := &jamfpro.ResourceMobileExtensionAttribute{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		DataType:         d.Get("data_type").(string),
		InventoryDisplay: d.Get("inventory_display").(string),
	}

	// Handle nested "input_type" field
	if v, ok := d.GetOk("input_type"); ok && len(v.([]interface{})) > 0 {
		inputTypeData := v.([]interface{})[0].(map[string]interface{})
		inputType := jamfpro.MobileExtensionAttributeSubsetInputType{
			Type:     inputTypeData["type"].(string),
		}

		

		attribute.InputType = inputType
	}

	// Serialize and pretty-print the Computer Extension Attribute object as XML for logging
	resourceXML, err := xml.MarshalIndent(attribute, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Jamf Pro Mobile Extension Attribute '%s' to XML: %v", attribute.Name, err)
	}

	log.Printf("[DEBUG] Constructed Jamf Pro Mobile Extension Attribute XML:\n%s\n", string(resourceXML))

	return attribute, nil
}
