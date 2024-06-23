// smartcomputergroup_object.go
package smartcomputergroups

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/endpoints/common/sharedschemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// constructJamfProSmartComputerGroup constructs a ResourceComputerGroup object from the provided schema data.
func constructJamfProSmartComputerGroup(d *schema.ResourceData) (*jamfpro.ResourceComputerGroup, error) {
	group := &jamfpro.ResourceComputerGroup{
		Name:    d.Get("name").(string),
		IsSmart: true,
	}

	group.Site = sharedschemas.ConstructSharedResourceSite(d.Get("site_id").(int))

	if v, ok := d.GetOk("criteria"); ok {
		group.Criteria = constructComputerGroupSubsetContainerCriteria(v.([]interface{}))
	}

	resourceXML, err := xml.MarshalIndent(group, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Jamf Pro Computer Group '%s' to XML: %v", group.Name, err)
	}

	log.Printf("[DEBUG] Constructed Jamf Pro Computer Group XML:\n%s\n", string(resourceXML))

	return group, nil
}

// constructComputerGroupSubsetContainerCriteria constructs a ComputerGroupSubsetContainerCriteria object from the provided schema data.
func constructComputerGroupSubsetContainerCriteria(criteriaList []interface{}) *jamfpro.ComputerGroupSubsetContainerCriteria {
	criteria := &jamfpro.ComputerGroupSubsetContainerCriteria{
		Size:      len(criteriaList),
		Criterion: &[]jamfpro.SharedSubsetCriteria{},
	}

	for _, item := range criteriaList {
		criterionData := item.(map[string]interface{})
		criterion := jamfpro.SharedSubsetCriteria{
			Name:         criterionData["name"].(string),
			Priority:     criterionData["priority"].(int),
			AndOr:        criterionData["and_or"].(string),
			SearchType:   criterionData["search_type"].(string),
			Value:        criterionData["value"].(string),
			OpeningParen: criterionData["opening_paren"].(bool),
			ClosingParen: criterionData["closing_paren"].(bool),
		}
		*criteria.Criterion = append(*criteria.Criterion, criterion)
	}

	return criteria
}
