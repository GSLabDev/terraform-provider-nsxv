package nsx

import (
	"strconv"
)

//get secuirty group details api
func SecurityGroupDetailsAPI(nsxCredentials NsxCredentials) string {
	securitygroupdetails := "https://" + nsxCredentials.ServerIP + ":" + strconv.Itoa(nsxCredentials.Port) + "/api/2.0/services/securitygroup/scope/globalroot-0"
	return securitygroupdetails
}

//add vm to security group api
func SecurityGroupAddMembersAPI(nsxCredentials NsxCredentials, securityGroupID string, virtualMachineID string) string {
	securityaddmembers := "https://" + nsxCredentials.ServerIP + ":" + strconv.Itoa(nsxCredentials.Port) + "/api/2.0/services/securitygroup/" + securityGroupID + "/members/" + virtualMachineID + "?failIfExists=true"
	return securityaddmembers
}

//check if NSX URL is valid or not url
func ConnectNSXAPI(serverIP string, port int) string {
	connect := "https://" + serverIP + ":" + string(port) + "/api"
	return connect
}

//remove vm from security group api
func RemoveVirtualMachineAPI(nsxCredentials NsxCredentials, securityGroupID string, virtualMachine string) string {
	delete := "https://" + nsxCredentials.ServerIP + ":" + strconv.Itoa(nsxCredentials.Port) + "/api/2.0/services/securitygroup/" + securityGroupID + "/members/" + virtualMachine + "?failIfAbsent=true"
	return delete
}

//read members in security group
func GetVirtualMachineInSecGroupAPI(nsxCredentials NsxCredentials, virtualMachine string) string {
	members := "https://" + nsxCredentials.ServerIP + ":" + strconv.Itoa(nsxCredentials.Port) + "/api/2.0/services/securitygroup/lookup/virtualmachine/" + virtualMachine
	return members
}
