package nsx

import (
	"encoding/xml"
	"fmt"
	"log"
	"reflect"
)

type securityGroupsMemberList struct {
	SecurityGroups securityGroups `xml:"securityGroups"`
}
type securityGroups struct {
	SecurityGroup []securityGroup `xml:"securitygroup"`
}

func getMembers(ResponseData []uint8, securityGroupName string) bool {
	var foundGroup = false
	//query the security group that has specified virtual machine
	var memberListQuery securityGroupsMemberList
	//unmarshal the response
	xml.Unmarshal([]byte(ResponseData), &memberListQuery)
	s := securityGroupsMemberList{}
	if reflect.DeepEqual(s, memberListQuery) != true {

		//search and log the security group members name
		for _, securityGroupMembers := range memberListQuery.SecurityGroups.SecurityGroup {
			if securityGroupName == securityGroupMembers.Name {
				log.Println("[INFO] Security Group found . ID " + securityGroupMembers.ObjectId + " Name " + securityGroupMembers.Name)
				foundGroup = true
				break
			} //if
		} //for
		return foundGroup
	} else {
		fmt.Errorf("[ERROR] No security group was found for specified virtual machine.")
		return foundGroup
	}
}

func checkMemberDestroyed(ResponseData []uint8, securityGroupName string) bool {
	//query the security group that has specified virtual machine
	var memberListQuery securityGroupsMemberList
	//unmarshal the response
	xml.Unmarshal([]byte(ResponseData), &memberListQuery)
	s := securityGroupsMemberList{}
	if reflect.DeepEqual(s, memberListQuery) != true {

		//search and log the security group members name
		for _, securityGroupMembers := range memberListQuery.SecurityGroups.SecurityGroup {
			if securityGroupMembers.Name == securityGroupName {
				return true
			}
		}
	}
	return false
}
