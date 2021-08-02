# Create a virtual network within the resource group
resource "azurerm_virtual_network" "pulsar" {
  name                = var.vnet_name
  resource_group_name = azurerm_resource_group.pulsar.name
  location            = azurerm_resource_group.pulsar.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "pulsar" {
  name                 = "sn-aks-pulsar"
  resource_group_name  = azurerm_resource_group.pulsar.name
  virtual_network_name = azurerm_virtual_network.pulsar.name
  address_prefixes     = ["10.0.1.0/24"]
}