package main

import "fmt"

type Student struct {
	name string
}

type String string

func (s String) append(str string) string {
	return fmt.Sprintf("%s%s", s, str)
}

// Define Method
func (s Student) getName() string {
	return s.name
}

// Value receiver
func (s Student) changeName() {
	// s là instance của struct Student
	s.name = "Robin"
	// fmt.Printf("%p\n", &s)
}

// Pointer receiver
func (s *Student) changeName2() {
	// s là pointer trỏ đến instance của struct Student
	s.name = "Robin"
	// fmt.Printf("%p\n", s)
}

func main() {
	student := &Student{"Ryan"}
	fmt.Printf("%p\n", student)
	fmt.Println(student)
	name := student.getName()

	fmt.Println(name)

	// Trong func changeName, s được tạo ra 1 giá trị mới nên khi thay đổi giá
	// trị của s sẽ không làm thay đổi giá trị của student.name. Trong func changeName,
	// in ra giá trị pointer để kiểm tra, giá trị pointer này sẽ khác với giá trị pointer
	// của biến student
	student.changeName()

	fmt.Println(student.name)

	// Trong func changeName2, s chính là giá trị student nên khi thay đổi giá trị
	// s thì student.name thay đổi theo, có thể in giá trị pointer ra để kiểm tra
	student.changeName2()
	fmt.Println(student.name)

	s1 := String("a")
	newStr := s1.append("b")
	fmt.Println(newStr)
}
