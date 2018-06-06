# nsx\_addvirtualmachinesecuritygroup

Provides a VMware vSphere's NSX variable description when one or more virtual machines are to be added to security group. 

## Example Usage

``` virtual_machine {
    name = "VM1"
    id   = "vm-40"
  }

```

## Argument Reference

The following arguments are supported:

* `name` - Provides virtual machine name that is to be added to the security group. 
* `id` - Provides ID of the corresponding virtual machine.
