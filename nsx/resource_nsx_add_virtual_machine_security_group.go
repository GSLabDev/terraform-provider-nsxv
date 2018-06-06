package nsx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

//structure of VirtualMachine details
type virtualMachineDetails struct {
	virtualMachinename string
	virtualMachineid   string
	clustername        string
	domainid           string
}

//function assigning structure of VirtualMachine details
func assignVirtualMachineDetails(virtualMachinename string, virtualMachineid string, clustername string, domainid string) virtualMachineDetails {
	var virtualMachineDetails virtualMachineDetails
	virtualMachineDetails.virtualMachinename = virtualMachinename
	virtualMachineDetails.virtualMachineid = virtualMachineid
	virtualMachineDetails.clustername = clustername
	virtualMachineDetails.domainid = domainid
	return virtualMachineDetails
} //assignDetails

func resourceNsxAddVirtualMachineSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxAddVirtualMachineCreate,
		Read:   resourceNsxAddVirtualMachineRead,
		Update: resourceNsxAddVirtualMachineUpdate,
		Delete: resourceNsxAddVirtualMachineDelete,
		Schema: map[string]*schema.Schema{

			"cluster_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Get Cluster name",
				ForceNew:    true,
			},

			"security_group_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Get security group name",
				ForceNew:    true,
			},

			"domain_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Get Cluster domain ID",
				ForceNew:    true,
			},

			"virtual_machine": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Set: resourceNsxAddVirtualMachineHash,
			},
		}, //schema
	}
} //resourceNSXSecurityGroup
//read the members (VirtualMachines) in security group
func resourceNsxAddVirtualMachineRead(d *schema.ResourceData, metadata interface{}) error {
	nsxCredentials := metadata.(NsxCredentials)
	securityGroupName := d.Get("security_group_name").(string)

	virtualMachine := d.Get("virtual_machine").(*schema.Set)
	for _, vms := range virtualMachine.List() {
		vm := vms.(map[string]interface{})

		//acqurie VirtualMachine details
		virtualMachineDetails := assignVirtualMachineDetails(vm["name"].(string), vm["id"].(string), d.Get("cluster_name").(string), d.Get("domain_id").(string))
		//Invoke the API responsible for listing the security groups that has specified VirtualMachine as a member
		requiredUrl := GetVirtualMachineInSecGroupAPI(nsxCredentials, virtualMachineDetails.virtualMachineid)
		//make a GET request

		responseBody, err := nsxCredentials.NsxConnection(GET, requiredUrl, nil)
		if err != nil {
			log.Printf("[ERROR] Virtual machine %s with ID %s not found or doesn't exists", vm["name"].(string), vm["id"].(string))
		}
		if responseBody != nil {
			//close the response body
			defer responseBody.Body.Close()
			//read the response body
			responseData, err := ioutil.ReadAll(responseBody.Body)
			if err != nil {
				log.Println(err)

			}
			//print the security group names into log) file that has the specified VirtualMachine as a member
			memberFound := getMembers(responseData, securityGroupName)
			//if members are not found set the partial ID for such VM as null
			if memberFound != true {
				virtualMachine.Remove(vms)
			}
		}
		readResourceSetPartial(d, virtualMachine)

	} //for

	if len(virtualMachine.List()) <= 0 {
		d.SetId("")
	}

	return nil
} //resourceNSXAddVirtualMachine

//add VirtualMachines to the security group
func resourceNsxAddVirtualMachineCreate(d *schema.ResourceData, metadata interface{}) error {
	virtualMachine := d.Get("virtual_machine").(*schema.Set)
	addVirtualMachines(d, metadata, virtualMachine)
	return nil
} //resourceNSXAddVirtualMachineCreate

func resourceNsxAddVirtualMachineDelete(d *schema.ResourceData, metadata interface{}) error {
	virtualMachine := d.Get("virtual_machine").(*schema.Set)
	removeVirtualMachines(d, metadata, virtualMachine)
	return nil
}

//update the partially added virtual machine
func resourceNsxAddVirtualMachineUpdate(d *schema.ResourceData, metadata interface{}) error {
	virtualMachine := d.Get("virtual_machine").(*schema.Set)

	if d.HasChange("virtual_machine") {
		oldList, newList := d.GetChange("virtual_machine")
		if oldList == nil {
			oldList = new(*schema.Set)
		} //if
		if newList == nil {
			newList = new(*schema.Set)
		} //if

		oList := oldList.(*schema.Set)
		nList := newList.(*schema.Set)

		toRemove := oList.Difference(nList)
		err := removeVirtualMachines(d, metadata, toRemove)
		if err != nil {
			log.Println(err)
		}

		toAdd := nList.Difference(oList)
		err = addVirtualMachines(d, metadata, toAdd)
		if err != nil {
			log.Println(err)
		}

		updateResourceSetPartial(d, virtualMachine)
	} //if

	return resourceNsxAddVirtualMachineRead(d, metadata)
} //update

func resourceNsxAddVirtualMachineHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["name"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["id"].(string)))

	return hashcode.String(buf.String())
}

func removeVirtualMachines(d *schema.ResourceData, metadata interface{}, toRemove *schema.Set) error {
	nsxCredentials := metadata.(NsxCredentials)
	securityGroupName := d.Get("security_group_name").(string)
	//get security group details i.e. id , node id , vsmuuid, revision number and description.
	securityGroupDetails := GetSecurityDetails(securityGroupName, nsxCredentials)

	for _, vms := range toRemove.List() {
		vm := vms.(map[string]interface{})

		err := resourceNsxAddVirtualMachineRead(d, metadata)
		if d.Id() == "" {
			return fmt.Errorf("[ERROR] Virtual Machine does not exists %s", err)
		}
		//acqurie VirtualMachine details
		virtualMachineId := vm["id"].(string)

		//add virtualMachine to the security group
		requiredUrl := RemoveVirtualMachineAPI(nsxCredentials, securityGroupDetails.ObjectIdDetail, virtualMachineId)
		//send a DELETE request
		response, err := nsxCredentials.NsxConnection(DELETE, requiredUrl, nil)
		if err != nil {
			log.Println(err)

		}

		//close the response body
		defer response.Body.Close()
	} //for
	return nil
} //removeNsxVirtualMachine

func addVirtualMachines(d *schema.ResourceData, metadata interface{}, toAdd *schema.Set) error {

	nsxCredentials := metadata.(NsxCredentials)
	securityGroupName := d.Get("security_group_name").(string)

	//get security group details i.e. id , node id , vsmuuid, revision number and description.
	securityGroupDetails := GetSecurityDetails(securityGroupName, nsxCredentials)
	for _, vms := range toAdd.List() {
		vm := vms.(map[string]interface{})

		//acqurie VirtualMachine details
		virtualMachineDetails := assignVirtualMachineDetails(vm["name"].(string), vm["id"].(string), d.Get("cluster_name").(string), d.Get("domain_id").(string))
		//add virtualMachine to the security group
		requiredUrl := SecurityGroupAddMembersAPI(nsxCredentials, securityGroupDetails.ObjectIdDetail, virtualMachineDetails.virtualMachineid)
		//get xml request body that is to be parsed
		data := parseXMLMarshal(securityGroupDetails, virtualMachineDetails)

		//send a PUT request
		response, err := nsxCredentials.NsxConnection(PUT, requiredUrl, strings.NewReader(data))

		if err != nil {
			log.Println(err)
			toAdd.Remove(vms)
		}
		if response != nil {
			//close the response body
			defer response.Body.Close()
			createResourceSetPartial(d, securityGroupDetails, securityGroupName)
		} //if

	} //for
	return nil
}

func createResourceSetPartial(d *schema.ResourceData, securityGroupDetails securityGroupDetails, securityGroupName string) {
	d.Partial(true)
	d.SetPartial("domain_id")
	d.SetPartial("security_group_name")
	d.SetPartial("cluster_name")
	d.SetPartial("virtual_machine")
	d.Partial(false)
	//set the id of the completed option to maintain the output,resources and primary id in the tfstate file useful at the time of terraform destroy
	d.SetId(d.Get("domain_id").(string) + "/" + securityGroupDetails.ObjectIdDetail + "/" + securityGroupName)

} //createResourceSetPartial

func updateResourceSetPartial(d *schema.ResourceData, virtualMachine *schema.Set) {
	d.Partial(true)
	d.Set("virtual_machine", virtualMachine)
	d.SetPartial("domain_id")
	d.SetPartial("security_group_name")
	d.SetPartial("cluster_name")
	d.SetPartial("virtual_machine")
	d.Partial(false)
}

func readResourceSetPartial(d *schema.ResourceData, virtualMachine *schema.Set) {
	d.Partial(true)
	d.Set("virtual_machine", virtualMachine)
	d.SetPartial("virtual_machine")
	d.Partial(false)
}
