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
    debug = true
  }

resource "dcloudtb_topology" "test_topology" {
  name = "VM Resource Test"
  description = "Testing Topology VM Resource Management"
  notes = "Created via Terraform Test"
  datacenter = "LON"
}

resource "dcloudtb_network" "routed_network" {
  name = "A routed network"
  description = "A Network to attach VMs"
  inventory_network_id = "L3-VLAN-2"
  topology_uid = dcloudtb_topology.test_topology.id
}

#data "dcloudtb_inventory_vms" "test_topology_inventory_vms" {
#  topology_uid = dcloudtb_topology.test_topology.id
#}
#
#output "inventory_vms" {
#  value = data.dcloudtb_inventory_vms.test_topology_inventory_vms
#}

resource "dcloudtb_vm" "vm1" {
  inventory_vm_id = "47161"
  topology_uid = dcloudtb_topology.test_topology.id
  name = "VM Created from Terraform"
  description = "It's Alive"
  cpu_qty = 2
  memory_mb = 4096
  nested_hypervisor = false
  os_family = "LINUX"

  advanced_settings {
    all_disks_non_persistent = false
    bios_uuid = "42 3a 5f 9d f1 a8 7c 0e-7d c2 44 27 2e d6 67 aa"
    name_in_hypervisor = "cmm"
    not_started = false
  }

  network_interfaces {
    network_uid = dcloudtb_network.routed_network.id
    name = "Network adapter 0"
    mac_address = "00:50:56:00:01:AA"
    type = "VIRTUAL_E1000"
  }

  network_interfaces {
    network_uid = dcloudtb_network.routed_network.id
    name = "Network adapter 1"
    mac_address = "00:50:56:00:01:AB"
    type = "VIRTUAL_E1000"
    ip_address = "127.0.0.2"
    ssh_enabled = true
    rdp_enabled = true
    rdp_auto_login = true
  }

  remote_access {
    username = "user"
    password = "password"
    vm_console_enabled = true

    display_credentials {
      username = "displayuser"
      password = "displaypassword"
    }

#    internal_urls {
#      location = "https://microsoft.com"
#      description = "microsoft"
#    }
#
#    internal_urls {
#      location = "https://google.com"
#      description = "google"
#    }
  }

  guest_automation {
    command = "RUN PROGRAM"
    delay_seconds = 10
  }
}

resource "dcloudtb_vm" "vm2" {
  inventory_vm_id = "47161"
  topology_uid = dcloudtb_topology.test_topology.id
  name = "2nd VM Created from Terraform"
  description = "It's Alive"
  cpu_qty = 2
  memory_mb = 3072
  nested_hypervisor = false
  os_family = "LINUX"

  advanced_settings {
    all_disks_non_persistent = false
    bios_uuid = "42 3a 5f 9d f1 a8 7c 0e-7d c2 44 27 2e d6 67 bb"
    name_in_hypervisor = "computer"
    not_started = false
  }

  network_interfaces {
    network_uid = dcloudtb_network.routed_network.id
    name = "Network adapter 0"
    mac_address = "00:50:56:00:02:AA"
    type = "VIRTUAL_E1000"
  }

  network_interfaces {
    network_uid = dcloudtb_network.routed_network.id
    name = "Network adapter 1"
    mac_address = "00:50:56:00:02:AB"
    type = "VIRTUAL_E1000"
    ip_address = "127.0.0.3"
    ssh_enabled = true
    rdp_enabled = true
    rdp_auto_login = true
  }

  remote_access {
    username = "user"
    password = "password"
    vm_console_enabled = true

    display_credentials {
      username = "displayuser"
      password = "displaypassword"
    }

#    internal_urls {
#      location = "https://cisco.com"
#      description = "cisco"
#    }
  }

  guest_automation {
    command = "RUN COMMAND"
    delay_seconds = 15
  }
}

output "vm" {
  value = dcloudtb_vm.vm2
}