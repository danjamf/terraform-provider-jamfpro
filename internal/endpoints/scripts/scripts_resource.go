// scripts_resource.go
package scripts

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceJamfProScripts defines the schema and CRUD operations for managing Jamf Pro Scripts in Terraform.
func ResourceJamfProScripts() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceJamfProScriptsCreate,
		ReadContext:   ResourceJamfProScriptsRead,
		UpdateContext: ResourceJamfProScriptsUpdate,
		DeleteContext: ResourceJamfProScriptsDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Read:   schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
			Delete: schema.DefaultTimeout(30 * time.Second),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Jamf Pro unique identifier (ID) of the script.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display name for the script.",
			},
			"category_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of the category to add the script to.",
			},
			"category_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Jamf Pro unique identifier (ID) of the category.",
			},
			"info": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Information to display to the administrator when the script is run.",
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Notes to display about the script (e.g., who created it and when it was created).",
			},
			"os_requirements": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The script can only be run on computers with these operating system versions. Each version must be separated by a comma (e.g., 10.11, 15, 16.1).",
			},
			"priority": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Execution priority of the script (BEFORE, AFTER, AT_REBOOT).",
				ValidateFunc: validation.StringInSlice([]string{"BEFORE", "AFTER", "AT_REBOOT"}, false),
			},
			"script_contents": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Contents of the script. Must be non-compiled and in an accepted format.",
			},
			"parameter4": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 4",
			},
			"parameter5": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 5",
			},
			"parameter6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 6",
			},
			"parameter7": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 7",
			},
			"parameter8": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 8",
			},
			"parameter9": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 9",
			},
			"parameter10": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 10",
			},
			"parameter11": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script parameter label 11",
			},
		},
	}
}

// ResourceJamfProScriptsCreate is responsible for creating a new Jamf Pro Script in the remote system.
// The function:
// 1. Constructs the script data using the provided Terraform configuration.
// 2. Calls the API to create the script in Jamf Pro.
// 3. Updates the Terraform state with the ID of the newly created script.
// 4. Initiates a read operation to synchronize the Terraform state with the actual state in Jamf Pro.
func ResourceJamfProScriptsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Assert the meta interface to the expected APIClient type
	apiclient, ok := meta.(*client.APIClient)
	if !ok {
		return diag.Errorf("error asserting meta as *client.APIClient")
	}
	conn := apiclient.Conn

	// Initialize variables
	var diags diag.Diagnostics

	// Construct the resource object
	resource, err := constructJamfProScript(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to construct Jamf Pro Script: %v", err))
	}

	// Retry the API call to create the resource in Jamf Pro
	var creationResponse *jamfpro.ResponseScriptCreate
	err = retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var apiErr error
		creationResponse, apiErr = conn.CreateScript(resource)
		if apiErr != nil {
			return retry.RetryableError(apiErr)
		}
		// No error, exit the retry loop
		return nil
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create Jamf Pro Script '%s' after retries: %v", resource.Name, err))
	}

	// Set the resource ID in Terraform state
	d.SetId(creationResponse.ID)

	// Retry reading the resource to ensure the Terraform state is up to date
	err = retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		readDiags := ResourceJamfProScriptsRead(ctx, d, meta)
		if len(readDiags) > 0 {
			return retry.RetryableError(fmt.Errorf(readDiags[0].Summary))
		}
		// Successfully read the resource, exit the retry loop
		return nil
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to synchronize Terraform state for Jamf Pro Script '%s' after creation: %v", resource.Name, err))
	}

	return diags
}

// ResourceJamfProScriptsRead is responsible for reading the current state of a Jamf Pro Script Resource from the remote system.
// The function:
// 1. Fetches the script's current state using its ID. If it fails then obtain script's current state using its Name.
// 2. Updates the Terraform state with the fetched data to ensure it accurately reflects the current state in Jamf Pro.
// 3. Handles any discrepancies, such as the script being deleted outside of Terraform, to keep the Terraform state synchronized.
func ResourceJamfProScriptsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Initialize API client
	apiclient, ok := meta.(*client.APIClient)
	if !ok {
		return diag.Errorf("error asserting meta as *client.APIClient")
	}
	conn := apiclient.Conn

	// Initialize variables
	var diags diag.Diagnostics
	resourceID := d.Id()

	// Read operation with retry
	var resource *jamfpro.ResourceScript

	err := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var apiErr error
		resource, apiErr = conn.GetScriptByID(resourceID)
		if apiErr != nil {
			if strings.Contains(apiErr.Error(), "404") || strings.Contains(apiErr.Error(), "410") {
				// Resource not found or gone, remove from Terraform state
				return retry.NonRetryableError(fmt.Errorf("resource not found, marked for deletion"))
			}
			// Convert any other API error into a retryable error to continue retrying
			return retry.RetryableError(apiErr)
		}
		// Successfully read the resource, exit the retry loop
		return nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "resource not found, marked for deletion") {
			d.SetId("") // Remove from Terraform state
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Resource not found or gone",
				Detail:   fmt.Sprintf("Jamf Pro Script with ID '%s' was not found on the server and is marked for deletion from terraform state.", resourceID),
			})
			return diags
		}
		// Handle the final error after all retries have been exhausted
		return diag.FromErr(fmt.Errorf("failed to read Jamf Pro Script with ID '%s' after retries: %v", resourceID, err))
	}

	// Update Terraform state with the resource information
	if err := d.Set("id", resource.ID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("name", resource.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("category_name", resource.CategoryName); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("category_id", resource.CategoryId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("info", resource.Info); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("notes", resource.Notes); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("os_requirements", resource.OSRequirements); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("priority", resource.Priority); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("script_contents", resource.ScriptContents); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter4", resource.Parameter4); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter5", resource.Parameter5); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter6", resource.Parameter6); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter7", resource.Parameter7); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter8", resource.Parameter8); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter9", resource.Parameter9); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter10", resource.Parameter10); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("parameter11", resource.Parameter11); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

// ResourceJamfProScriptsUpdate is responsible for updating an existing Jamf Pro Department on the remote system.
func ResourceJamfProScriptsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Initialize API client
	apiclient, ok := meta.(*client.APIClient)
	if !ok {
		return diag.Errorf("error asserting meta as *client.APIClient")
	}
	conn := apiclient.Conn

	// Initialize variables
	var diags diag.Diagnostics
	resourceID := d.Id()

	// Construct the resource object
	resource, err := constructJamfProScript(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to construct Jamf Pro Script for update: %v", err))
	}

	// Update operation with retries
	err = retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
		_, apiErr := conn.UpdateScriptByID(resourceID, resource)
		if apiErr != nil {
			// If updating by ID fails, attempt to update by Name
			resourceName := d.Get("name").(string)
			_, apiErrByName := conn.UpdateScriptByName(resourceName, resource)
			if apiErrByName != nil {
				// If updating by name also fails, return a retryable error
				return retry.RetryableError(apiErrByName)
			}
		}
		// Successfully updated the resource, exit the retry loop
		return nil
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update Jamf Pro Script '%s' (ID: %s) after retries: %v", d.Get("name").(string), resourceID, err))
	}

	// Read the resource to ensure the Terraform state is up to date
	readDiags := ResourceJamfProScriptsRead(ctx, d, meta)
	if len(readDiags) > 0 {
		diags = append(diags, readDiags...)
	}

	return diags
}

// ResourceJamfProScriptsDelete is responsible for deleting a Jamf Pro Department.
func ResourceJamfProScriptsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Initialize API client
	apiclient, ok := meta.(*client.APIClient)
	if !ok {
		return diag.Errorf("error asserting meta as *client.APIClient")
	}
	conn := apiclient.Conn

	// Initialize variables
	var diags diag.Diagnostics
	resourceID := d.Id()

	// Use the retry function for the delete operation with appropriate timeout
	err := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		// Attempt to delete the resource by ID
		apiErr := conn.DeleteScriptByID(resourceID)
		if apiErr != nil {
			// If deleting by ID fails, attempt to delete by Name
			resourceName := d.Get("name").(string)
			apiErrByName := conn.DeleteScriptByName(resourceName)
			if apiErrByName != nil {
				// If deletion by name also fails, return a retryable error
				return retry.RetryableError(apiErrByName)
			}
		}
		// Successfully deleted the resource, exit the retry loop
		return nil
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete Jamf Pro Script '%s' (ID: %s) after retries: %v", d.Get("name").(string), resourceID, err))
	}

	// Clear the ID from the Terraform state as the resource has been deleted
	d.SetId("")

	return diags
}
