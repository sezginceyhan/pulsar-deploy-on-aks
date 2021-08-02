variable "resource_group" {
  type = string
}

variable "aks_name" {
  type = string
}

variable "aks_agent_count" {
  default = 1
}

variable "vnet_name" {
  type = string
}