package controllers

import (
	"github.com/wy3148/autoptesting/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/wy3148/autoptesting/store"
)


type ErrInfo struct{
	Error string `json:"error"`
	Message string `json:"message"`
}



type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	cid := c.GetString(":contact_id")
	beego.Debug("Getting contact by id:",cid)
	if cid != "" {
		contact, err := store.Store().GetContact(cid)
		if contact == nil || err != nil{
			beego.Debug("contact is not found in local store")
			contact, err = getRemoteContact(cid)
			if contact == nil || err != nil{
				beego.Error("failed to get contact from remote server:",err)
			}else{
				store.Store().UpdateContact(contact)
				c.Data["json"] = contact
			}
		}else{
			beego.Debug("contact is retrieved from local store")
			c.Data["json"] = contact
		}
	}

	//if contact is not found, we still call this fun, it return {}
	c.ServeJSON()
}

func (c *ContactController) Post() {
	var contact models.NewContactReq
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &contact)
	if err != nil{
		beego.Error("failed to validate the contact infomation:",err)
		c.Data["json"] = &ErrInfo{
			Error:"invalid contact",
			Message:"failed to parse the request",
		}
		c.ServeJSON()
		return
	}

	if contact.Req.Email == ""{
		c.Data["json"] = &ErrInfo{
			Error:"invalid contact",
			Message:"email can not be empty",
		}
		c.ServeJSON()
		return
	}


	//we firstly update contact to remote server, then override the local cache
	updated , err := updateRmoteContact(&contact)
	if err != nil{
		c.Data["json"] = &ErrInfo{
			Error:"failed to update contact",
			Message:err.Error(),
		}
		c.ServeJSON()
		return
	}
	contact.Req.Id = updated
	err = store.Store().UpdateContact(contact.Req)
	if err != nil{
		beego.Warn("failed to update contact into cache system",err)
	}
	c.Data["json"] = contact.Req
	c.ServeJSON()
}


