terraform {
  required_providers {
    dcloudtb = {
      version = "0.1"
      source  = "cisco.com/dcloud/dcloudtb"
    }
  }
}

provider "dcloudtb" {
  tb_url = "https://tbv3-dev.dev.ciscodcloud.com/api"
}

resource "dcloudtb_topology" "test_topology" {
  name        = "Test Topology For Testing Inventory Networks"
  description = "Will be used to load inventory networks"
  notes       = ""
  datacenter  = "LON"
}

data "dcloudtb_inventory_networks" "topology1_inventory_networks" {
  topology_uid = dcloudtb_topology.test_topology.id
}

output "networks" {
  value = data.dcloudtb_inventory_networks.topology1_inventory_networks
}