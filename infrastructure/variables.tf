variable "tenancy_ocid" {}

variable "user_ocid" {}

variable "private_key_path" {}

variable "fingerprint" {}

variable "region" {}

variable "compartment_ocid" {}

variable "ssh_authorized_keys" {}

variable "availability_domain" {}

variable "instance_name" {}

variable "subnet_id" {}

variable "ipv6_cidr" {}

variable "operating_system" {
  type = string
  default = "Oracle Linux"
}

variable "operating_system_version" {
  type = string
  default = "9"
}
