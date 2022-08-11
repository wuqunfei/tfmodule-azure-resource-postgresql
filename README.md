## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | 3.17.0 |
| <a name="provider_random"></a> [random](#provider\_random) | 3.3.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_postgresql_database.postgresqldb](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_database) | resource |
| [azurerm_postgresql_server.postgresqlserver](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_server) | resource |
| [azurerm_resource_group.rg](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/resource_group) | resource |
| [random_password.password](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | Name for the resource group holding resources for this example | `string` | `"terratest-postgres-rg"` | no |
| <a name="input_location"></a> [location](#input\_location) | The Azure region in which to deploy this sample | `string` | `"West Europe"` | no |
| <a name="input_postfix"></a> [postfix](#input\_postfix) | A postfix string to centrally mitigate resource name collisions. | `string` | `"resource"` | no |
| <a name="input_db_user"></a> [db\_user](#input\_db\_user) | Database User Name | `string` | `"pgadmin"` | no |
| <a name="input_db_password"></a> [db\_password](#input\_db\_password) | Database User Password | `string` | `"2021#Munich"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_sku_name"></a> [sku\_name](#output\_sku\_name) | n/a |
| <a name="output_servername"></a> [servername](#output\_servername) | n/a |
| <a name="output_rgname"></a> [rgname](#output\_rgname) | n/a |
