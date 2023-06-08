terraform {
  required_providers {
    dcloudtb = {
      version = "0.1"
      source  = "cisco.com/dcloud/dcloudtb"
    }
  }
}

provider "dcloudtb" {
  tb_url = "https://tbv3-production.ciscodcloud.com/api"
}

resource "dcloudtb_topology" "test_topology" {
  name        = "HW Resource Test"
  description = "Testing Topology HW Resource Management"
  notes       = "Created via Terraform Test"
  datacenter  = "LON"
}

resource "dcloudtb_network" "routed_network" {
  name                 = "A routed network"
  description          = "Demonstrating a network routed through VPOD Gateway"
  inventory_network_id = "L3-VLAN-2"
  topology_uid         = dcloudtb_topology.test_topology.id
}

#data "dcloudtb_inventory_hws" "test_topology_inventory_hws" {
#  topology_uid = dcloudtb_topology.test_topology.id
#}
#
#output "hws" {
#  value = data.dcloudtb_inventory_hws.test_topology_inventory_hws
#}

#data "dcloudtb_inventory_hw_scripts" "test_topology_inventory_hw_scripts" {
#  topology_uid = dcloudtb_topology.test_topology.id
#}
#
#output "scripts" {
#  value = data.dcloudtb_inventory_hw_scripts.topology1_inventory_hw_scripts
#}

#data "dcloudtb_inventory_hw_template_configs" "test_topology_inventory_hw_template_configs" {
#  topology_uid = dcloudtb_topology.test_topology.id
#}
#
#output "configs" {
#  value = data.dcloudtb_inventory_hw_template_configs.test_topology_inventory_hw_template_configs
#}

resource "dcloudtb_hw" "cx_core2" {
  topology_uid               = dcloudtb_topology.test_topology.id
  inventory_hw_id            = "170"
  name                       = "CX Core2 Device"
  hardware_console_enabled   = true
  power_control_enabled      = false
  startup_script_uid         = "bjlfkxev55nh35eh6kku13971"
  custom_script_uid          = "668eljku7jwpk8bpysz5njyrz"
  shutdown_script_uid        = "435ya6tjh5u4uv3ku2kphesr"
  template_config_script_uid = "79ila00mn7icfbtk3dg7fuasy"

  network_interfaces {
    network_interface_id = "GigabitEthernet1/0/24"
    network_uid          = dcloudtb_network.routed_network.id
  }
}

#data "dcloudtb_hws" "test_topology_hws" {
#  topology_uid = dcloudtb_topology.test_topology.id
#}
#
#output "topology_hws" {
#  value = data.dcloudtb_hws.test_topology_hws
#}