package store

import (
	"github.com/wy3148/autoptesting/models"
	"testing"
	"time"
)

func TestStoreContact(t *testing.T){

	c := &models.Contact{
		Email:         "test1234@gmail.com",
		Created:       time.Now(),
		Update:		   time.Now(),
		ApiOriginated: true,
		CustomFiled:   []models.Field{
			{
					Kind:"Test Field",
					Val:"This is a test",
					FieldType:"string",
					Deleted:false,
			},
		},
		Name:          "test1234 test1234",
		LastName:      "test1234",
		FirstName:     "test1234",
		Id:            "person_test1234",
	}

	s := Store()

	err := s.UpdateContact(c)
	if err != nil{
		t.Fatal(err)
	}

	contactInCache, err := s.GetContact(c.Id)
	if err != nil{
		t.Fatal(err)
	}

	if c.Email != contactInCache.Email ||
		c.Id != contactInCache.Id{
		t.Error("contact information is inconsistent in store")
	}
}