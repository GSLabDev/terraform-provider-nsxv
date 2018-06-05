
# nsx\_addvirtualmachinesecuritygroup

Provides a VMware vSphere's NSX resource. This can be used to create(add) and delete Virtual Machine in an existing on NSX Server security group. It adds a virtual machine using NSX REST API. 

## Example Usage

```hcl
resource "nsx_add_virtual_machine_security_group" "virtual_machine" {
  count                = "${length(var.virtual_machine_name_list)}"
  cluster_name         = "Compute Cluster A"
  security_group_name  = "Security group 1"
  domain_id            = "domain-c242"
  virtual_machine_name = "${element(var.virtual_machine_name_list,count.index)}"
  virtual_machine_id   = "${element(var.virtual_machine_id_list,count.index)}"
}

```

## Argument Reference

The following arguments are supported:

* `cluster_name` - (Required) The cluster name on which virtual machine exists
* `security_group_name` - (Required) Add virtual machine to specified security group
* `domain_id` - (Required) The domain ID of the cluster on which virtual machine exist
* `virtual_machine_name` - (Required) The virtual machine name
* `virtual_machine_id` - (Required) Respective virtual machine id