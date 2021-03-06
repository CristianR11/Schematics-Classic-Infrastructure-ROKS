resource "ibm_container_cluster" "cluster" {
  name              = "${var.cluster_name}${random_id.name.hex}"
  datacenter        = "${var.datacenter}"
  default_pool_size = 1
  machine_type      = "${var.machine_type}"
  hardware          = "${var.hardware}"
  kube_version      = "${var.kube_version}"
  public_vlan_id    = "${var.public_vlan_id}"
  private_vlan_id   = "${var.private_vlan_id}"
  lifecycle {
  ignore_changes = ["kube_version"]
  }
  resource_group_id = "8f8d7c37ad0c4e29912f7057ea7f90cf"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = "${ibm_container_cluster.cluster.id}"
}

resource "random_id" "name" {
  byte_length = 4
}
