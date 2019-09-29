package models

import "time"

//
//{
//    "Email": "test11@slarty.com",
//    "created_at": "2019-09-28T14:14:29.000Z",
//    "updated_at": "2019-09-28T14:16:06.000Z",
//    "api_originated": true,
//    "custom_fields": [
//        {
//            "kind": "Test Field",
//            "value": "This is a test",
//            "fieldType": "string",
//            "deleted": false
//        }
//    ],
//    "Name": "Slarty1 Bartfast2",
//    "LastName": "Bartfast22",
//    "FirstName": "Slarty12",
//    "contact_id": "person_447AA9AE-B311-45E7-9FF2-E846FA5D939E"
//}

type Field struct{
	Kind string `json:"kind"`
	Val string 	`json:"value"`
	FieldType string  `json:"fieldType"`
	Deleted bool `json:"deleted"`
}

type Contact struct {
	Email 	string `json:"Email"`
	Created 	time.Time `json:"created_at,omitempty"`
	Update 	time.Time `json:"updated_at,omitempty"`
	ApiOriginated bool `json:"api_originated,omitempty"`
	CustomFiled []Field `json:"custom_fields,omitempty"`
	Name string `json:"Name",omitempty`
	LastName string `json:"LastName,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	Id string `json:"contact_id,omitempty"`
}

type NewContactReq struct{
	Req *Contact `json:"contact"`
}