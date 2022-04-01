package test

import (
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func Test002(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../modules/test_002_add_custom_core",
		PlanFilePath: "./tfplan",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Get the plan and store as a struct in memory
	ps := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

	keys := make([]string, 0, len(ps.ResourcePlannedValuesMap))
	mgs := []string{"module.test_core.azurerm_management_group.level_1[\"/providers/Microsoft.Management/managementGroups/1245\"]"}

	// Iterate over every planned value and output the name
	for k := range ps.ResourcePlannedValuesMap {
		keys = append(keys, k)
	}

	// Log to act as a breakpoint
	log.Print(ps.RawPlan.TerraformVersion)

	require.Subset(
		t,
		keys,
		mgs,
		"But something...")
}
