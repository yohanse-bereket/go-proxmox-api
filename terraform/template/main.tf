resource "proxmox_lxc" "container" {

  hostname = var.container_name
  target_node = "pve"

  cores  = var.cpu
  memory = var.memory

  ostemplate = "local:vztmpl/ubuntu-22.04-standard_22.04-1_amd64.tar.zst"

  password = "changeme"

  rootfs {
    storage = "local-lvm"
    size    = "8G"
  }

  network {
    name   = "eth0"
    bridge = "vmbr0"
    ip     = "dhcp"
  }

  start = true
}