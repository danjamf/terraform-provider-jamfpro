// mobileextensionattributes_data_validation.go
package mobileextensionattributes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// validateJamfProRResourceMobileExtensionAttributesDataFields performs custom validation on the Resource's schema so that passed values from
// teraform resource declarations align with attibute combinations supported by the Jamf Pro api.
func validateJamfProRResourceMobileExtensionAttributesDataFields(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error {
	// Extract the first item from the input_type list, which should be a map
	inputTypes, ok := diff.GetOk("input_type")
	if !ok || len(inputTypes.([]interface{})) == 0 {
		return fmt.Errorf("input_type must be provided")
	}

	inputTypeMap := inputTypes.([]interface{})[0].(map[string]interface{})

	inputType := inputTypeMap["type"].(string)
	
	choices := inputTypeMap["choices"].([]interface{})

	switch inputType {
	case "Pop-up Menu":
		// Ensure "choices" is populated
		if len(choices) == 0 {
			return fmt.Errorf("'choices' must be populated when input_type is 'Pop-up Menu'")
		}
	case "Text Field":
		// Ensure neither "script", "platform" nor "choices" are populated
		if len(choices) > 0 {
			return fmt.Errorf("'choices' must not be populated when input_type is 'Text Field'")
		}
	}

	return nil
}
