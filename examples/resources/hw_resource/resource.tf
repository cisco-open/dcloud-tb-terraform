terraform {
  required_providers {
    dcloud = {
      version = "0.1"
      source  = "cisco-open/dcloud"
    }
  }
}

provider "dcloud" {
  tb_url = "https://tbv3-production.ciscodcloud.com/api"
}

resource "dcloud_topology" "test_topology" {
  name        = "HW Resource Test"
  description = "Testing Topology HW Resource Management"
  notes       = "Created via Terraform Test"
  datacenter  = "LON"
}

resource "dcloud_network" "routed_network" {
  name                 = "A routed network"
  description          = "Demonstrating a network routed through VPOD Gateway"
  inventory_network_id = "L3-VLAN-2"
  topology_uid         = dcloud_topology.test_topology.id
}

resource "dcloud_hw" "IE4000" {
  topology_uid               = dcloud_topology.test_topology.id
  inventory_hw_id            = "76"
  name                       = "IE 4000 Device"
  hardware_console_enabled   = false
  startup_script_uid         = "bjlfkxev55nh35eh6kku13971"
  custom_script_uid          = "668eljku7jwpk8bpysz5njyrz"
  shutdown_script_uid        = "435ya6tjh5u4uv3ku2kphesr"
  template_config_script_uid = "79ila00mn7icfbtk3dg7fuasy"

  network_interfaces {
    network_interface_id = "GigabitEthernet1/0/24"
    network_uid          = dcloud_network.routed_network.id
  }
}