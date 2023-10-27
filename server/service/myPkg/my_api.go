package myPkg

import "fmt"

type MyApi struct {
}

func (m *MyApi) CreateApiS() {
	fmt.Println("Hello world!")
}
