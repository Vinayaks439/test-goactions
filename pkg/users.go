package pkg

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUser(name string, age int) {
	user := User{}
	user.Name = name
	user.Age = age
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("User Created: ", string(b))
}

func DeleteUser(name string, age int) {
	user := User{}
	user.Name = name
	user.Age = age
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("User Deleted: ", string(b))
}
