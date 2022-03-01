package main

import "fmt"

type Inter interface {
	changeName(name string)
}

type Student struct {
	Id    string
	Name  string
	Birth int
}

func (st *Student) changeName(name string) {
	st.Name = name
}

func main() {
	var student1 Inter = &Student{"1", "Linh", 1991}
	student1.changeName("a")
	fmt.Println(student1)
}
