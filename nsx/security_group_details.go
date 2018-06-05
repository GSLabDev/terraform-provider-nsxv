package nsx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//structure of main tag list and inner tag security group
type list struct {
	SecurityGroupInfo []securityGroup `xml:"securitygroup"`
}

//security group details
type securityGroup struct {
	Name        string `xml:"name"`
	ObjectId    string `xml:"objectId"`
	VsmUuid     string `xml:"vsmUuid"`
	NodeId      string `xml:"nodeId"`
	Revision    int    `xml:"revision"`
	Description string `xml:"description"`
}

//structure assigning security details
type securityGroupDetails struct {
	NameDetail        string
	DescriptionDetail string
	ObjectIdDetail    string
	VsmUuidDetail     string
	NodeIdDetail      string
	RevisionDetail    int
}

func GetSecurityDetails(securityGroupName string, nsxCredentials NsxCredentials) securityGroupDetails {
	requiredUrl := SecurityGroupDetailsAPI(nsxCredentials)
	//send a request
	response, err := nsxCredentials.NsxConnection(GET, requiredUrl, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	//read the response body
	body, err := ioutil.ReadAll(response.Body) //read the response (in xml format)

	//read the xml and markup the security group details
	var list list
	var securityGroupDetails securityGroupDetails

	xml.Unmarshal(body, &list) //extract the xml format
	//iterate through the security group details and find the expected security group
	for _, group := range list.SecurityGroupInfo {
		if group.Name == securityGroupName { //check for the specified secuirty group
			securityGroupDetails.NameDetail = group.Name               //asign security group name
			securityGroupDetails.NodeIdDetail = group.NodeId           //assign security group node ID
			securityGroupDetails.VsmUuidDetail = group.VsmUuid         //assign security group VsmUuid
			securityGroupDetails.RevisionDetail = group.Revision       //assign security group revision number
			securityGroupDetails.ObjectIdDetail = group.ObjectId       //assign security group ID
			securityGroupDetails.DescriptionDetail = group.Description //assign security group descdription if available
		} //if
	} //for
	return securityGroupDetails
} //getSecurityDetails
