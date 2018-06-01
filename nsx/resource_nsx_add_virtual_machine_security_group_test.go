package nsx

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAddVirtualMachineSecurityGroup_Basic(t *testing.T) {
	virtualMachineName := "VM1"
	resourceName := "nsx_add_virtual_machine_security_group.virtual_machine"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { TestNsxPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAddVirtualMachineSecurityGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAddVirtualMachineSecurityGroupConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAddVirtualMachineSecurityGroupExists(resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "virtual_machine_name", virtualMachineName),
				),
			},
		},
	})
}
func testAddVirtualMachineSecurityGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Virtual Machine is added to security group")
		}

		virtualMachineId := rs.Primary.Attributes["virtual_machine_id"]

		nsxCredentials := testAccProvider.Meta().(NsxCredentials)
		requiredUrl := GetVirtualMachineInSecGroupAPI(nsxCredentials, virtualMachineId)

		responseBody, err := nsxCredentials.NsxConnection(GET, requiredUrl, nil)
		if err != nil {
			return err
		}
		defer responseBody.Body.Close()

		responseData, err := ioutil.ReadAll(responseBody.Body)
		if err != nil {
			return err
		}

		memberFound := getMembers(responseData)
		if memberFound != true {
			return fmt.Errorf("[ERROR] No Security Group found for virtual machine ")
		}
		return nil
	}
}

func testAddVirtualMachineSecurityGroupDestroy(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: " + n)
		}

		virtualMachineId := rs.Primary.Attributes["virtual_machine_id"]
		securityGroupName := rs.Primary.Attributes["security_group_name"]

		nsxCredentials := testAccProvider.Meta().(NsxCredentials)
		requiredUrl := GetVirtualMachineInSecGroupAPI(nsxCredentials, virtualMachineId)

		responseBody, err := nsxCredentials.NsxConnection(GET, requiredUrl, nil)
		if err != nil {
			return err
		}
		defer responseBody.Body.Close()

		responseData, err := ioutil.ReadAll(responseBody.Body)
		if err != nil {
			return err
		}

		destroyStatus := checkMemberDestroyed(responseData, securityGroupName)
		if destroyStatus {
			return fmt.Errorf("[ERROR] Virtual machine not deleted from security group")
		}
		return nil
	}
} //func testAddVirtualMachineSecurityGroupDestroy

func testAddVirtualMachineSecurityGroupConfigBasic() string {
	return fmt.Sprintf(`
		resource "nsx_add_virtual_machine_security_group" "virtual_machine" {
			cluster_name         = "%s"
			security_group_name  = "%s"
			domain_id            = "%s"
			virtual_machine_name = "%s"
			virtual_machine_id   = "%s"
		}`, os.Getenv("NSX_CLUSTER_NAME"),
		os.Getenv("SECURITY_GROUP_NAME"),
		os.Getenv("DOMAIN_ID"),
		os.Getenv("VIRTUAL_MACHINE_NAME"),
		os.Getenv("VIRTUAL_MACHINE_ID"))
}
