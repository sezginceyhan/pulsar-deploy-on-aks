terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "2.68.0"
    }
  }
}

provider "azurerm" {
  # Configuration options
  features {}
}


resource "azurerm_resource_group" "pulsar" {
  name     = var.resource_group
  location = "West Europe"
  tags = {
    Environment         = "experimental"
    CanBeDeletedAnytime = "True"
  }
}

resource "azurerm_kubernetes_cluster" "pulsar" {
  name                = var.aks_name
  location            = azurerm_resource_group.pulsar.location
  resource_group_name = azurerm_resource_group.pulsar.name
  dns_prefix          = replace(var.aks_name, "([^\\w\\d\\s])", "")

  default_node_pool {
    name       = "default"
    node_count = var.aks_agent_count
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment         = "experimental"
    CanBeDeletedAnytime = "True"
  }
}

resource "local_file" "client_certificate" {
  content         = azurerm_kubernetes_cluster.pulsar.kube_config.0.client_certificate
  filename        = "./.kube/${var.aks_name}.crt"
  file_permission = "0600"
}

resource "local_file" "kube_config" {
  content         = azurerm_kubernetes_cluster.pulsar.kube_config_raw
  filename        = "./.kube/config"
  file_permission = "0600"
}



