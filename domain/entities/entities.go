package entities

type User struct {
	Id   string ` ksql:"id" json:"id"`
	Name string `ksql:"name" json:"name"`
}

type Address struct {
	Id     string ` ksql:"id" json:"id"`
	Street string `ksql:"street" json:"street"`
	Name   string `ksql:"name" json:"name"`
	UserId string `ksql:"user_id" json:"userId"`
}
