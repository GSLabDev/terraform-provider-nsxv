


# Terraform Vmware NSX Provider

This is the repository for the Terraform [VMware NSX][1] Provider, which one can use
with Terraform to work with VMware NSX.

[1]: https://www.vmware.com/in/products/nsx.html

Coverage is currently only limited to add Virtual Machines to security group in VMware's vSphere NSX.
Watch this space!

For general information about Terraform, visit the [official website][3] and the
[GitHub project page][4].

[3]: https://terraform.io/
[4]: https://github.com/hashicorp/terraform

# Using the Provider

The current version of this provider requires Terraform v0.11.7 or higher to
run.

Note that you need to run `terraform init` to fetch the provider before
deploying. Read about the provider split and other changes to TF v0.11.7 in the
official release announcement found [here][4].

[4]: https://www.hashicorp.com/blog/hashicorp-terraform-0-10/

## Full Provider Documentation

The provider is useful in adding Virtual Machines to security group in VMware's NSX.

### Example
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
## Checking the Logs
To persist logged output you can set TF_LOG_PATH in order to force the log to always be appended to a specific file when logging is enabled. Note that even when TF_LOG_PATH is set, TF_LOG must be set in order for any logging to be enabled.

To check logs use the following commands
```sh
export TF_LOG=DEBUG
export TF_LOG_PATH=/home/terraform-provider-nsx/logs/nsx.log
```

# Building The Provider

**NOTE:** Unless you are [developing][7] or require a pre-release bugfix or feature,
you will want to use the officially released version of the provider (see [the
section above][8]).

[7]: #developing-the-provider
[8]: #using-the-provider


## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/terraform-providers/terraform-provider-nsx`:

```sh
mkdir -p $GOPATH/src/github.com/terraform-providers
cd $GOPATH/src/github.com/terraform-providers
git clone git@github.com:terraform-providers/terraform-provider-nsx
```

## Running the Build

After the clone has been completed, you can enter the provider directory and
build the provider.

```sh
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-nsx
make build
```

## Installing the Local Plugin

After the build is complete, copy the `terraform-provider-nsx` binary into
the same path as your `terraform` binary, and re-run `terraform init`.

After this, your project-local `.terraform/plugins/ARCH/lock.json` (where `ARCH`
matches the architecture of your machine) file should contain a SHA256 sum that
matches the local plugin. Run `shasum -a 256` on the binary to verify the values
match.

# Developing the Provider

If you wish to work on the provider, you'll first need [Go][9] installed on your
machine (version 1.9+ is **required**). You'll also need to correctly setup a
[GOPATH][10], as well as adding `$GOPATH/bin` to your `$PATH`.

[9]: https://golang.org/
[10]: http://golang.org/doc/code.html#GOPATH

See [Building the Provider][11] for details on building the provider.

[11]: #building-the-provider

## Configuring Environment Variables

Most of the tests in this provider require a comprehensive list of environment
variables to run. See the individual `*_test.go` files in the
[`nsx/`](nsx/) directory for more details. The next section also
describes how you can manage a configuration file of the test environment
variables.

### Using the `.tf-nsx-devrc.mk` file

The [`tf-nsx-devrc.mk`](tf-nsx-devrc.mk) file contains
an up-to-date list of environment variables required to run the acceptance
tests. Copy this to `$HOME/.tf-nsx-devrc.mk` and change the permissions to
something more secure (ie: `chmod 600 $HOME/.tf-nsx-devrc.mk`), and
configure the variables accordingly.

## Running the Acceptance Tests

After this is done, you can run the acceptance tests by running:

```sh
$ make testacc
```

If you want to run against a specific set of tests, run `make testacc` with the
`TESTARGS` parameter containing the run mask as per below:

```sh
make testacc TESTARGS="-run=TestAccNsx"
```

This following example would run all of the acceptance tests matching
`TestAccNsx`. Change this for the specific tests you want to
run.

