# Connect-AzAccount -EnvironmentName AzureUSGovernment

Add-AzAccount
#Select the correct subscription
Get-AzSubscription -SubscriptionName "AzIntConsumption" | Select-AzSubscription

terraform init -upgrade

terraform plan

terraform apply -auto-approve -parallelism 50    # (default 20 upper limit 50)