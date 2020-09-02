package repository

import (
	"context"
	"jwt-client/config"

	"gopkg.in/mgo.v2/bson"
)

type repo struct {
	DB         string
	Collection string
}

type CRUD interface {
	Get(bson.M) (interface{}, error)
	Create(interface{}) (interface{}, error)
	Update(bson.M, bson.M) (interface{}, error)
	Delete(bson.M) (interface{}, error)
}

func NewCRUD(db, col string) CRUD {
	return &repo{DB: db, Collection: col}
}

func (re *repo) Get(f bson.M) (interface{}, error) {
	datas := make(map[string]interface{})
	var data []map[string]interface{}

	c, e := config.GetMongoClient()
	if e != nil {
		return nil, e
	}
	defer c.Disconnect(context.Background())

	r, e := c.Database(re.DB).Collection(re.Collection).Find(context.Background(), f)
	if e != nil {
		return nil, e
	}

	if e = r.All(context.Background(), &data); e != nil {
		return nil, e
	}

	datas["data"] = data
	datas["total_data"] = len(data)
	return datas, nil
}

func (re *repo) Create(p interface{}) (interface{}, error) {
	c, e := config.GetMongoClient()
	if e != nil {
		return nil, e
	}
	defer c.Disconnect(context.Background())

	r, e := c.Database(re.DB).Collection(re.Collection).InsertOne(context.Background(), p)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (re *repo) Update(f, u bson.M) (interface{}, error) {
	c, e := config.GetMongoClient()
	if e != nil {
		return nil, e
	}
	defer c.Disconnect(context.Background())

	r, e := c.Database(re.DB).Collection(re.Collection).UpdateOne(context.Background(), f, u)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (re *repo) Delete(f bson.M) (interface{}, error) {
	c, e := config.GetMongoClient()
	if e != nil {
		return nil, e
	}
	defer c.Disconnect(context.Background())

	r, e := c.Database(re.DB).Collection(re.Collection).DeleteOne(context.Background(), f)
	if e != nil {
		return nil, e
	}

	return r, nil
}
