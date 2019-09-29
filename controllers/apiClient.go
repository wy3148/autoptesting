package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wy3148/autoptesting/models"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"github.com/astaxie/beego"
)


//autop simple http client

var apiKey string

const(
	remoteServer = "https://api2.autopilothq.com/v1/contact"
	apiHeader = "autopilotapikey"
)

func init(){
	apiKey = os.Getenv("autoptesting_key")
	if apiKey == ""{
		panic("api key is not configured, export autoptesting_key=api_key before run application")
	}
}

//Get contact from remote autopilot api service
func getRemoteContact(id string) (*models.Contact, error){
	beego.Debug("req url is:",remoteServer + "/" + id)
	req ,err := http.NewRequest("GET",remoteServer + "/" + id, nil)
	if err != nil {
		return nil,err
	}

	req.Header.Add(apiHeader,apiKey)
	hc := http.Client{
		Timeout: time.Duration(time.Second * 30),
	}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error code: %d, body:%s", resp.StatusCode, string(respBody))
	}

	var c models.Contact
	err = json.Unmarshal(respBody,&c)
	if err != nil{
		return nil,err
	}
	return &c, nil
}

func updateRmoteContact(c * models.NewContactReq)(string,error){
	beego.Debug("req url is:",remoteServer)
	b, err := json.Marshal(c)
	if err != nil{
		return "",err
	}

	beego.Debug("request body is ",string(b))

	req ,err := http.NewRequest("POST",remoteServer, bytes.NewReader(b))
	if err != nil {
		return "",err
	}

	//using postman, it's text/plain, body is json formatted string
	//accepete by api service
	req.Header.Add("Content-Type","text/plain")
	req.Header.Add(apiHeader,apiKey)
	hc := http.Client{
		Timeout: time.Duration(time.Second * 30),
	}
	resp, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Error code: %d, body:%s", resp.StatusCode, string(respBody))
	}

	var updated models.Contact
	err = json.Unmarshal(respBody,&updated)
	if err != nil{
		return "",err
	}
	return updated.Id, nil
}


