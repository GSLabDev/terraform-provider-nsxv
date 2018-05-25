package nsx

import (
	"encoding/xml"
	"log"
)

//Security group tag and subtag
type securityGroupList struct {
	ObjectId           string             `xml:"objectId"`
	VsmUuid            string             `xml:"vsmUuid"`
	NodeId             string             `xml:"nodeId"`
	Revision           int                `xml:"revision"`
	Name               string             `xml:"name"`
	Description        string             `xml:"description"`
	Member             member             `xml:"member"`
	TypeSecurity       typeSecurity       `xml:"type"`
	Scope              scope              `xml:"scope"`
	ExtendedAttributes extendedAttributes `xml:"extendedAttributes"`
	ClientHandle       string             `xml:"clientHandle"`
}

//structure of scope tag
type scope struct {
	Id             string `xml:"id"`
	ObjectTypeName string `xml:"objectTypeName"`
	Name           string `xml:"name"`
}

//structure of member tag
type member struct {
	ObjectId           string       `xml:"objectId"`
	ObjectTypeName     string       `xml:"objectTypeName"`
	VsmUuid            string       `xml:"vsmUuid"`
	NodeId             string       `xml:"nodeId"`
	Revision           int          `xml:"revision"`
	Name               string       `xml:"name"`
	ClientHandle       string       `xml:"clientHandle"`
	ExtendedAttributes string       `xml:"extendedAttributes"`
	IsUniversal        bool         `xml:"isUniversal"`
	UniversalRevision  int          `xml:"universalRevision"`
	TypeSecurity       typeSecurity `xml:"type"`
	Scope              scope        `xml:"scope"`
}

//structure of extendedAttributes tag
type extendedAttributes struct {
	ExtendedAttribute []extendedAttribute `xml:"extendedAttribute"`
}

//structure of extendedattribute subtag
type extendedAttribute struct {
	Name  string `xml:"name"`
	Value bool   `xml:"value"`
}

//structure of type tag
type typeSecurity struct {
	TypeName string `xml:"typeName"`
}

//parse the xml formed
func parseXMLMarshal(securityGroupDetails securityGroupDetails, virtualMachineDetails virtualMachineDetails) string {

	requestBody := securityGroupList{
		//assigning security group details
		ObjectId: securityGroupDetails.ObjectIdDetail,
		VsmUuid:  securityGroupDetails.VsmUuidDetail,
		NodeId:   securityGroupDetails.NodeIdDetail,
		Revision: securityGroupDetails.RevisionDetail,
		TypeSecurity: typeSecurity{
			TypeName: "SecurityGroup",
		},
		Scope: scope{
			Id:             "globalroot-0",
			ObjectTypeName: "GlobalRoot",
			Name:           "Global",
		},
		ExtendedAttributes: extendedAttributes{
			ExtendedAttribute: []extendedAttribute{
				{
					Name:  "localMembersOnly",
					Value: false,
				},
			},
		},
		//defining the type of the members to be added into security group i.e. Virtual Machine
		Member: member{
			ObjectId:       virtualMachineDetails.virtualMachineid,
			ObjectTypeName: "VirtualMachine",
			VsmUuid:        securityGroupDetails.VsmUuidDetail,
			NodeId:         securityGroupDetails.NodeIdDetail,
			Revision:       securityGroupDetails.RevisionDetail,
			Name:           virtualMachineDetails.virtualMachinename,
			TypeSecurity: typeSecurity{
				TypeName: "VirtualMachine",
			},

			Scope: scope{
				Id:             virtualMachineDetails.domainid,
				ObjectTypeName: "ClusterComputeResource",
				Name:           virtualMachineDetails.clustername,
			},
			IsUniversal:       false,
			UniversalRevision: 0,
		},
	}

	//formating the xml including the indentation
	securityDetailsBody, err := xml.MarshalIndent(&requestBody, "", "\t")
	if err != nil {
		log.Println(err)
	}
	//return the xml
	return string(securityDetailsBody)
}
