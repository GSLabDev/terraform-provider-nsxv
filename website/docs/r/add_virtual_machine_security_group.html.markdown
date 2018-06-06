
# nsx\_addvirtualmachinesecuritygroup

Provides a VMware vSphere's NSX resource. This can be used to create(add) and delete Virtual Machine in an existing on NSX Server security group. It adds a virtual machine using NSX REST API. 

## Example Usage

```#add virtual machines in the list to the specified security group
resource "nsx_add_virtual_machine_security_group" "virtualmachine" {
 
  cluster_name         = "Compute Cluster A"
  security_group_name  = "Security Group 1"
  domain_id            = "domain-c242"
 
  virtual_machine {
    name = "VM1"
    id   = "vm-40"
  }

  virtual_machine {
    name = "VM2"
    id   = "vm-41"
  }

  virtual_machine {
    name = "VM3"
    id   = "vm-42"
  }

  virtual_machine {
    name = "VM4"
    id   = "vm-56"
  }
}

```

## Argument Reference

The following arguments are supported:

* `cluster_name` - (Required) The cluster name on which virtual machine exists
* `security_group_name` - (Required) Add virtual machine to specified security group
* `domain_id` - (Required) The domain ID of the cluster on which virtual machine exist
* `virtual_machine` - (Required) The virtual machine parameter consists of virtual machine name and ID as a set
