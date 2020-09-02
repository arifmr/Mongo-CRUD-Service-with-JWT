package repository

import (
	"context"
	"jwt-client/config"
	"jwt-client/model"

	"gopkg.in/mgo.v2/bson"
)

type user struct {
	Username string
	Password string
}

type Auth interface {
	Login() (model.User, error)
	Logout() (map[string]interface{}, error)
}

func NewAuth() Auth {
	return &user{}
}

func (p *user) Login() (model.User, error) {
	var datas model.User

	c, e := config.GetMongoClient()
	if e != nil {
		return model.User{}, e
	}
	defer c.Disconnect(context.Background())

	f := bson.M{"$and": []bson.M{bson.M{"username": p.Username}, bson.M{"password": p.Password}}}
	e = c.Database("management").Collection("user").FindOne(context.Background(), f).Decode(&datas)
	if e != nil {
		return model.User{}, e
	}

	return datas, nil
}

func (p *user) Logout() (map[string]interface{}, error) {
	datas := make(map[string]interface{})

	c, e := config.GetMongoClient()
	if e != nil {
		return nil, e
	}
	defer c.Disconnect(context.Background())

	return datas, nil
}
