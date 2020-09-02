package service

import (
	"jwt-client/model"
	"jwt-client/repository"

	"gopkg.in/mgo.v2/bson"
)

func GetUserByID(id string) (interface{}, error) {
	f := bson.M{"_id": id}

	c := repository.NewCRUD("management", "user")
	r, e := c.Get(f)
	if e != nil {
		return nil, e
	}

	return r, e
}

func GetUser() (interface{}, error) {
	f := bson.M{}

	c := repository.NewCRUD("management", "user")
	r, e := c.Get(f)
	if e != nil {
		return nil, e
	}

	return r, e
}

func CreateUser(p model.User) (interface{}, error) {
	c := repository.NewCRUD("management", "user")
	r, e := c.Create(p)
	if e != nil {
		return nil, e
	}

	return r, e
}

func UpdateUser(p model.User, id string) (interface{}, error) {
	f := bson.M{"_id": id}
	u := bson.M{"$set": p}

	c := repository.NewCRUD("management", "user")
	r, e := c.Update(f, u)
	if e != nil {
		return nil, e
	}

	return r, e
}

func DeleteUser(id string) (interface{}, error) {
	f := bson.M{"_id": id}

	c := repository.NewCRUD("management", "user")
	r, e := c.Delete(f)
	if e != nil {
		return nil, e
	}

	return r, e
}
