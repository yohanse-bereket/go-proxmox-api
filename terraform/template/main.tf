resource "proxmox_virtual_environment_container" "container" {
  node_name = "pve"
unprivileged  = true
  initialization {
    hostname = var.container_name

    ip_config {
      ipv4 {
        address = "dhcp"
      }
    }

    user_account {
      password = "newpass"
    }
  }

  cpu {
    cores = var.cpu
  }

  memory {
    dedicated = var.memory
  }

  operating_system {
    template_file_id = "local:vztmpl/ubuntu-24.04-standard_24.04-2_amd64.tar.zst"
    type             = "ubuntu"
  }

  disk {
    datastore_id = "local-lvm"
    size         = 8
  }

  network_interface {
   name   = "eth0"
   bridge = "vmbr0"
  }

  started = true
}
