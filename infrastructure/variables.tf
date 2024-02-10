variable "tenancy_ocid" {}

variable "user_ocid" {}

variable "private_key_path" {}

variable "fingerprint" {}

variable "region" {}

variable "compartment_ocid" {}

variable "ssh_authorized_keys" {}

variable "subnet_id" {}

variable "ipv6_cidr" {}

variable "availability_domain" {
  type    = string
  default = "ldMg:EU-FRANKFURT-1-AD-1"
}

variable "instance_name" {
  type    = string
  default = "1.cloud"
}

variable "operating_system" {
  type    = string
  default = "Oracle Linux"
}

variable "operating_system_version" {
  type    = string
  default = "9"
}
