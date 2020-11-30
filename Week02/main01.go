package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := Biz()
	fmt.Printf("service: %+v\n", err)
}

type user struct {
	id   int
	name string
	age  int
}

var ErrNotFound = errors.New("Record not found!")

func Dao() (error) {
	return errors.Wrap(ErrNotFound, "SQL failed query!")
}

var ErrNotWork = errors.New("Too young to work!")

func Biz() error {
	u := &user{id: 1, name: "Tom", age: 17}

	err := Dao()
	if err != nil {
		return err
	}

	if u.age < 18 {
		return errors.Wrap(ErrNotWork, "")
	}
	return nil
}
