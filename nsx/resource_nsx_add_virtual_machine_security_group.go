package nsx

import (
<<<<<<< HEAD
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
=======
	"fmt"
	"io/ioutil"
	"log"

	"strings"

>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5
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
<<<<<<< HEAD
=======
			"virtual_machine_name": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "Get virtual machine name",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"virtual_machine_id": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "Get virtual machine ID",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5

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
<<<<<<< HEAD

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
=======
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5
		}, //schema
	}
} //resourceNSXSecurityGroup
//read the members (VirtualMachines) in security group
func resourceNsxAddVirtualMachineRead(d *schema.ResourceData, metadata interface{}) error {
	nsxCredentials := metadata.(NsxCredentials)
<<<<<<< HEAD
	securityGroupName := d.Get("security_group_name").(string)

	virtualMachine := d.Get("virtual_machine").(*schema.Set)
	log.Println("vj----read", virtualMachine)
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
		log.Println("vj----read--after", virtualMachine, len(virtualMachine.List()))

		d.Partial(true)
		d.Set("virtual_machine", virtualMachine)
		d.SetPartial("virtual_machine")
		d.Partial(false)

	} //for

	if len(virtualMachine.List()) <= 0 {
		d.SetId("")
	}

=======
	virtualMachineNames := d.Get("virtual_machine_name").([]interface{})
	virtualMachineId := d.Get("virtual_machine_id").([]interface{})
	securityGroupName := d.Get("security_group_name").(string)
	var AddedVMID []string   //:= make([]string, len(virtualMachineId)) //d.Set("virtual_machine_id",virtualMachineId)
	var AddedVMName []string // := make([]string, len(virtualMachineNames))  //d.Set("virtual_machine_id",virtualMachineId)

	if len(virtualMachineNames) > 0 {

		machine := 0
		for machine < len(virtualMachineNames) && (len(virtualMachineNames) == len(virtualMachineId)) {

			//acqurie VirtualMachine details
			virtualMachineDetails := assignVirtualMachineDetails(virtualMachineNames[machine].(string), virtualMachineId[machine].(string), d.Get("cluster_name").(string), d.Get("domain_id").(string))
			//Invoke the API responsible for listing the security groups that has specified VirtualMachine as a member
			requiredUrl := GetVirtualMachineInSecGroupAPI(nsxCredentials, virtualMachineDetails.virtualMachineid)
			//make a GET request
			responseBody, err := nsxCredentials.NsxConnection(GET, requiredUrl, nil)
			if err != nil {
				log.Printf("[ERROR] Virtual machine with ID %s not found or doesn't exists", virtualMachineNames[machine].(string))
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
				if memberFound {
					AddedVMID = append(AddedVMID, virtualMachineId[machine].(string))
					AddedVMName = append(AddedVMName, virtualMachineNames[machine].(string))
				}
			}

			d.Partial(true)
			d.Set("virtual_machine_id", AddedVMID)
			d.Set("virtual_machine_name", AddedVMName)
			d.SetPartial("virtual_machine_name")
			d.SetPartial("virtual_machine_id")
			d.Partial(false)

			machine = machine + 1
		} //for

		if len(AddedVMID) <= 0 {
			d.SetId("")
		}

	} //if
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5
	return nil
} //resourceNSXAddVirtualMachine

//add VirtualMachines to the security group
func resourceNsxAddVirtualMachineCreate(d *schema.ResourceData, metadata interface{}) error {
<<<<<<< HEAD

	nsxCredentials := metadata.(NsxCredentials)
=======
	nsxCredentials := metadata.(NsxCredentials)

	virtualMachineNames := d.Get("virtual_machine_name").([]interface{})
	virtualMachineId := d.Get("virtual_machine_id").([]interface{})
	var SuccessAddVMID []string   //:= make([]string, len(virtualMachineId)) //d.Set("virtual_machine_id",virtualMachineId)
	var SuccessAddVMName []string // := make([]string, len(virtualMachineNames))  //d.Set("virtual_machine_id",virtualMachineId)

>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5
	securityGroupName := d.Get("security_group_name").(string)

	//get security group details i.e. id , node id , vsmuuid, revision number and description.
	securityGroupDetails := GetSecurityDetails(securityGroupName, nsxCredentials)
<<<<<<< HEAD
	virtualMachine := d.Get("virtual_machine").(*schema.Set)
	log.Println(reflect.TypeOf(virtualMachine))
	for _, vms := range virtualMachine.List() {
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
			virtualMachine.Remove(vms)
		}
		if response != nil {
			//close the response body
			defer response.Body.Close()
			d.Partial(true)
			d.SetPartial("domain_id")
			d.SetPartial("security_group_name")
			d.SetPartial("cluster_name")
			d.SetPartial("virtual_machine")
			d.Partial(false)
			//set the id of the completed option to maintain the output,resources and primary id in the tfstate file useful at the time of terraform destroy
			d.SetId(d.Get("domain_id").(string) + "/" + securityGroupDetails.ObjectIdDetail + "/" + securityGroupName)

		} //if

	} //for
=======

	if (len(virtualMachineNames)) > 0 && (len(virtualMachineNames) == len(virtualMachineId)) {

		machine := 0
		for machine < len(virtualMachineNames) {

			//acqurie VirtualMachine details
			virtualMachineDetails := assignVirtualMachineDetails(virtualMachineNames[machine].(string), virtualMachineId[machine].(string), d.Get("cluster_name").(string), d.Get("domain_id").(string))
			//add virtualMachine to the security group
			requiredUrl := SecurityGroupAddMembersAPI(nsxCredentials, securityGroupDetails.ObjectIdDetail, virtualMachineDetails.virtualMachineid)
			//get xml request body that is to be parsed
			data := parseXMLMarshal(securityGroupDetails, virtualMachineDetails)

			//send a PUT request
			response, err := nsxCredentials.NsxConnection(PUT, requiredUrl, strings.NewReader(data))

			if err != nil {
				log.Println(err)
			}
			if response != nil {
				//close the response body
				defer response.Body.Close()
				SuccessAddVMID = append(SuccessAddVMID, virtualMachineId[machine].(string))
				SuccessAddVMName = append(SuccessAddVMName, virtualMachineNames[machine].(string))

				d.Partial(true)
				d.Set("virtual_machine_id", SuccessAddVMID)
				d.Set("virtual_machine_name", SuccessAddVMName)
				d.SetPartial("domain_id")
				d.SetPartial("security_group_name")
				d.SetPartial("cluster_name")
				d.SetPartial("virtual_machine_id")
				d.SetPartial("virtual_machine_name")
				d.Partial(false)
				//set the id of the completed option to maintain the output,resources and primary id in the tfstate file useful at the time of terraform destroy
				d.SetId(d.Get("domain_id").(string) + "/" + securityGroupDetails.ObjectIdDetail + "/" + securityGroupName)

			} //if

			machine = machine + 1
		} //for

	} //if
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5

	return nil
} //resourceNSXAddVirtualMachineCreate

func resourceNsxAddVirtualMachineDelete(d *schema.ResourceData, metadata interface{}) error {
<<<<<<< HEAD
	nsxCredentials := metadata.(NsxCredentials)

	securityGroupName := d.Get("security_group_name").(string)
	//get security group details i.e. id , node id , vsmuuid, revision number and description.
	securityGroupDetails := GetSecurityDetails(securityGroupName, nsxCredentials)

	virtualMachine := d.Get("virtual_machine").(*schema.Set)
	for _, vms := range virtualMachine.List() {
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
			return err
		}

		//close the response body
		defer response.Body.Close()
	} //for

=======

	nsxCredentials := metadata.(NsxCredentials)

	securityGroupName := d.Get("security_group_name").(string)
	virtualMachineIdList := d.Get("virtual_machine_id").([]interface{})
	//get security group details i.e. id , node id , vsmuuid, revision number and description.
	securityGroupDetails := GetSecurityDetails(securityGroupName, nsxCredentials)

	if len(virtualMachineIdList) > 0 {

		machine := 0
		for machine < len(virtualMachineIdList) {

			err := resourceNsxAddVirtualMachineRead(d, metadata)
			if d.Id() == "" {
				return fmt.Errorf("[ERROR] Virtual Machine does not exists %s", err)
			}
			//acqurie VirtualMachine details
			virtualMachineId := virtualMachineIdList[machine].(string)

			//add virtualMachine to the security group
			requiredUrl := RemoveVirtualMachineAPI(nsxCredentials, securityGroupDetails.ObjectIdDetail, virtualMachineId)
			//send a DELETE request
			response, err := nsxCredentials.NsxConnection(DELETE, requiredUrl, nil)
			if err != nil {
				log.Println(err)
				return err
			}

			//close the response body
			defer response.Body.Close()
			machine = machine + 1
		} //for
	} //if
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5
	return nil
}

//update the partially added virtual machine
func resourceNsxAddVirtualMachineUpdate(d *schema.ResourceData, metadata interface{}) error {
<<<<<<< HEAD
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

		toRemove := oList.Difference(nList).List()
		log.Println("remove-------", toRemove)

		toAdd := nList.Difference(oList).List()
		log.Println("add------", toAdd)

		if len(toAdd) > 0 {
			virtualMachine.Add(toAdd)
		}

		if len(toRemove) > 0 {
			virtualMachine.Remove(toRemove)
		}

		d.Partial(true)
		d.Set("virtual_machine", virtualMachine)
		d.SetPartial("domain_id")
		d.SetPartial("security_group_name")
		d.SetPartial("cluster_name")
		d.SetPartial("virtual_machine")
=======
	if d.HasChange("virtual_machine_id") {
		oldList, newList := d.GetChange("virtual_machine_id")
		addVMs := getDifference(oldList.([]interface{}), newList.([]interface{}))
		d.Partial(true)
		d.Set("virtual_machine_id", addVMs)
		d.SetPartial("virtual_machine_id")
		d.Partial(false)

	} //if

	if d.HasChange("virtual_machine_name") {
		oldList, newList := d.GetChange("virtual_machine_name")
		addVMs := getDifference(oldList.([]interface{}), newList.([]interface{}))
		d.Partial(true)
		d.Set("virtual_machine_name", addVMs)
		d.SetPartial("virtual_machine_name")
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5
		d.Partial(false)

	} //if

	return resourceNsxAddVirtualMachineRead(d, metadata)
} //update

<<<<<<< HEAD
func resourceNsxAddVirtualMachineHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["name"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["id"].(string)))

	return hashcode.String(buf.String())
=======
func getDifference(oldList []interface{}, newList []interface{}) interface{} {
	var diff []interface{}
	for _, i := range oldList {
		for _, j := range newList {
			if i == j {
				diff = append(diff, i.(string))
			}
		} //for j
	} //for i
	return diff
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5
}
