# NSX Provider

The NSX provider is used to interact with the resources supported by VMware NSX REST API.
The provider needs to be configured with the proper credentials before it can be used.

~> **NOTE:** The provider at this time only supports adding Virtual Machines to security group.

## Example Usage

```nsx
# Configure the NSX Provider
provider "nsx" {
  nsx_username  = "${var.nsx_username}"
  nsx_password  = "${var.nsx_password}"
  nsx_server_ip = "${var.nsx_server_ip}"
  port          = ${var.nsx_port}
}

#add virtual machines in the list to the specified security group
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

The following arguments are used to configure the Active Directory Provider:

* `nsx_username` - (Required) This is the username for NSX server. Can also
  be specified with the `NSX_SERVER_USERNAME` environment variable.
* `nsx_password` - (Required) This is the password for NSX server. Can
  also be specified with the `NSX_SERVER_PASSWORD` environment variable.
* `nsx_server_ip` - (Required) This is the NSX server ip for executing REST API operations.
 Can also be specified with the `NSX_SERVER_IP ` environment  variable.
* `port` - (Required) This is the port for API operations of the NSX using 443.
Can also be specified with the `NSX_SERVER_PORT ` environment variable.

