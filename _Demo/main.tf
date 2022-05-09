# Get the current client configuration from the AzureRM provider.​

# This is used to populate the root_parent_id variable with the​

# current Tenant ID used as the ID for the "Tenant Root Group"​

# Management Group.​



data "azurerm_client_config" "core" {}

# Declare the Terraform Module for Cloud Adoption Framework​

# Enterprise-scale and provide a base configuration.​


module "enterprise_scale" {

  source = "Azure/caf-enterprise-scale/azurerm"

  version = "1.1.3"



  providers = {

    azurerm = azurerm

    azurerm.connectivity = azurerm.connectivity

    azurerm.management = azurerm.management

  }


  root_parent_id = data.azurerm_client_config.core.tenant_id

  root_id = var.root_id

  root_name = var.root_name


  deploy_core_landing_zones = false

}