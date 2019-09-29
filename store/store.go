package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"github.com/wy3148/autoptesting/models"
)

type StoreIf interface {
	GetContact(id string) (*models.Contact, error)
	UpdateContact(c *models.Contact) error
}

type storeDb struct {
	redisCli *redis.Client
}

var cache *storeDb

func Store() StoreIf {

	if cache != nil {
		return cache
	}

	//application can load conf/app.conf by default, but we add this for
	//unit testing purpose
	err := beego.LoadAppConfig("ini", os.Getenv("GOPATH")+"/src/github.com/wy3148/autoptesting/conf/app.conf")
	if err != nil {
		panic(err)
	}

	r := beego.AppConfig.String("redisUrl")
	if len(r) == 0 {
		panic("Redis server url is not configured")
	}

	client := redis.NewClient(&redis.Options{
		Addr: r,
	})
	cache = &storeDb{redisCli: client}
	return cache
}

func (s *storeDb) GetContact(id string) (*models.Contact, error) {
	contact, err := s.redisCli.Get(id).Result()
	if err != nil {
		return nil, err
	}

	var c models.Contact
	err = json.Unmarshal([]byte(contact), &c)

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *storeDb) UpdateContact(c *models.Contact) error {

	//email must be filled
	if len(c.Email) == 0 {
		return errors.New("invalid contact, email is missing")
	}

	if len(c.Id) == 0 {
		return errors.New("invalid contact, contact id is missing")
	}

	contact, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = s.redisCli.Set(c.Id, contact, 0).Err()
	if err != nil {
		//if we failed to update the contact, we can delete it
		//from cache system, this is to make sure we can always
		//get the latest contact from remote server
		s.redisCli.Del(c.Id)
		return err
	}
	return nil
}
