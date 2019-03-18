package main

import (
	"encoding/json"
	"fmt"

	"github.com/audio35444/kibanalogger/examples/domain"
)

func main() {
	e := domain.Entity{1234, "entity1", "this is an entity"}
	fmt.Println(e)

	p := domain.Person{"Juan", "Perez", int32(30)}
	fmt.Println(p)

	str, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(str))
}
