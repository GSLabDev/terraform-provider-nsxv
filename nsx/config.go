package nsx

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
)

//API methods

var GET = "GET"
var POST = "POST"
var PUT = "PUT"
var DELETE = "DELETE"
var OPTIONS = "OPTIONS"

//NSX login credentials structure
type NsxCredentials struct {
	ServerIP string
	Port     int
	Username string
	Password string
}

type Error struct {
	Details string `xml:"details"`
}

//check the nsx credentails
func NsxConfig(d *schema.ResourceData) (interface{}, error) {

	//acquire the nsxcredentials from provider and verify the nsx server api
	serverIP := d.Get("nsx_server_ip").(string)
	port := d.Get("port").(int)
	username := d.Get("nsx_username").(string)
	password := d.Get("nsx_password").(string)

	//check if server IP is not empty
	if serverIP == "" {
		return nil, fmt.Errorf("[ERROR] no nsx server IP was found.")
	}

	//check connection to nsx server assuming provided url is valid
	nsxUrl := ConnectNSXAPI(serverIP, port) //get url for connection
	_, err := http.NewRequest(OPTIONS, nsxUrl, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//create nsx credentials interface
	config := NsxCredentials{
		ServerIP: serverIP,
		Username: username,
		Password: password,
		Port:     port,
	}

	return config, nil
} //func nsxConfig

func (nsxCredentials NsxCredentials) NsxConnection(method string, url string, buffer io.Reader) (*http.Response, error) {
	// Initialize the HTTPS client to skip SSL certificate verification
	flagSsl := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	skipSslVerify := &http.Client{Transport: flagSsl}

	//make request to the REST API of NSX
	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Configure Basic authentication and Content-Type
	request.SetBasicAuth(nsxCredentials.Username, nsxCredentials.Password)
	request.Header.Set("Content-Type", "application/xml")

	//Make request to the NSX REST API
	response, err := skipSslVerify.Do(request)
	//check the status of the request i.e. response
	status := httpResponseStatus(response)

	if response.StatusCode == http.StatusOK {
		return response, nil
	}

<<<<<<< HEAD
	return nil, fmt.Errorf("[ERROR] %s", status)
=======
	return response, fmt.Errorf("[ERROR] %s", status)
>>>>>>> 1a71950fbb51dd501309515de24d21d2474570a5

} // nsxConnection

func httpResponseStatus(response *http.Response) string {
	var status string
	if response.StatusCode == http.StatusBadRequest {
		buffer, _ := ioutil.ReadAll(response.Body)
		log.Println(string(buffer))
		status = readErrorResponse(string(buffer))
	}
	if response.StatusCode == http.StatusNotFound {
		status = "Couldn't retrieve the content.404 Not found"
	}
	if response.StatusCode == http.StatusUnauthorized {
		status = "Invalid user login credentials. Please check username / password"
	}
	return status
}

func readErrorResponse(data string) string {
	var errMsg Error
	err := xml.Unmarshal([]byte(data), &errMsg)

	if err != nil {
		fmt.Println(err)
	}

	return errMsg.Details
}
