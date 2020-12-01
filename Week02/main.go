package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := Biz()
	if err != nil {
		// 捕获错误，进行相应处理
		fmt.Printf(err.Error())
	}
	// 输出错误及堆栈信息
	//fmt.Printf("service: %+v\n", err)
}

type user struct {
	id   int
	name string
	age  int
}

var ErrNotFound = errors.New("Record not found!")

func Dao() error {
	// 向上抛出错误
	return errors.Wrap(ErrNotFound, "SQL failed query!")
}

var ErrNotWork = errors.New("Too young to work!")

func Biz() error {
	u := &user{id: 1, name: "Tom", age: 17}

	err := Dao()
	if err != nil {
		// 向上抛出错误
		return err
	}

	if u.age < 18 {
		// 向上抛出错误
		return errors.Wrap(ErrNotWork, "")
	}
	return nil
}
