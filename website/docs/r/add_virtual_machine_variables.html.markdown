# nsx\_addvirtualmachinesecuritygroup

Provides a VMware vSphere's NSX variable description when one or more virtual machines are to be added to security group. 

## Example Usage

```hcl
variable "virtual_machine_name_list" {
  type    = "list"
  default = ["VM1", "VM2", "VM3"]
}

variable "virtual_machine_id_list" {
  type    = "list"
  default = ["vm-296", "vm-298", "vm-297"]
}
```

## Argument Reference

The following arguments are supported:

* `type` - Provide type as list of arguments 
* `default` - Provide list of arguments