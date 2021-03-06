variable "vpc_cidr" {
  default = "10.1.0.0/21"
}

variable "vswitch_cidr" {
  default = "10.1.0.0/24"
}

variable "entry_cidr" {
  default = "172.11.1.1/32"
}

variable "rule_policy" {
  default = "accept"
}

variable "instance_type" {
  default = "ecs.n4.small"
}

variable "image_id" {
  default = "ubuntu_140405_64_40G_cloudinit_20161115.vhd"
}

variable "internet_charge_type" {
  default = "PayByTraffic"
}

