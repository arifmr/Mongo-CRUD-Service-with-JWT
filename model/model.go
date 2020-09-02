package model

type User struct {
	Username string `json:"username" bson:"username" example:"ramailh"`
	Password string `json:"password" bson:"password" example:"secretpass"`
	Email    string `json:"email" bson:"email" example:"ilhamr174@gmail.com	"`
	FullName string `json:"full_name" bson:"full_name" example:"Ilham Ramadhan"`
}

type Person struct {
	Username string `json:"username" bson:"username" example:"ramailh"`
	Password string `json:"password" bson:"password" example:"secretpass"`
	Email    string `json:"email" bson:"email" example:"ilhamr174@gmail.com	"`
	FullName string `json:"full_name" bson:"full_name" example:"Ilham Ramadhan"`
}

type Client struct {
	Username string `json:"username" bson:"username" example:"ramailh"`
	Password string `json:"password" bson:"password" example:"secretpass"`
	Email    string `json:"email" bson:"email" example:"ilhamr174@gmail.com	"`
	FullName string `json:"full_name" bson:"full_name" example:"Ilham Ramadhan"`
}

type Response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}
