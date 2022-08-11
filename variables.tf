# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# ARM_CLIENT_ID
# ARM_CLIENT_SECRET
# ARM_SUBSCRIPTION_ID
# ARM_TENANT_ID

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "resource_group_name" {
  description = "Name for the resource group holding resources for this example"
  type        = string
  default     = "terratest-postgres-rg"
}

variable "location" {
  description = "The Azure region in which to deploy this sample"
  type        = string
  default     = "West Europe"
}

variable "postfix" {
  description = "A postfix string to centrally mitigate resource name collisions."
  type        = string
  default     = "resource"
}

variable "db_user" {
  description = "Database User Name"
  type        = string
  default     = "pgadmin"
}

variable "db_password" {
  description = "Database User Password"
  type        = string
  default     = "2021#Munich"
}