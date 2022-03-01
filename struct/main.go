package main

import "fmt"

type StudentInfo struct {
	class string
	email string
	age   int
}

type Student struct {
	id   int
	name string
}

type Student2 struct {
	id   int
	name string
	info StudentInfo
}

func main() {
	st1 := Student{
		id:   123,
		name: "Ryan",
	}

	p3 := &st1
	p3.id = 1234

	fmt.Println(p3)
	fmt.Printf("%+v\n", st1)

	fmt.Println(st1)
	fmt.Println(st1.id)

	st2 := Student{123, "Robin"}
	fmt.Println(st2)

	var st3 Student = struct {
		id   int
		name string
	}{
		id:   777,
		name: "Bill",
	}

	fmt.Println(st3)

	// annomous struct
	var anonymous = struct {
		email string
		age   int
	}{
		email: "ryan@gmail.com",
		age:   27,
	}
	fmt.Println(anonymous)

	// pointer tro toi struct
	p5 := Student{
		456,
		"Robin",
	}
	fmt.Println(p5)
	fmt.Printf("%T", p5)
	fmt.Println()
	fmt.Println("p5", &p5)

	fmt.Println(p5.id)
	fmt.Println(p5.name)

	// anonymous fields
	type NoName struct {
		string
		int
	}

	n := NoName{"Nguyen Van A", 6666}
	fmt.Println(n)

	student := Student2{
		id:   123,
		name: "Ryan",
		info: StudentInfo{
			class: "ql091",
			email: "ryan@gmail.com",
			age:   27,
		},
	}

	fmt.Println(student)

	// compare 2 struct
	type struct1 struct {
		id   int
		name string
	}

	s1 := struct1{id: 1, name: "A"}
	s2 := struct1{id: 1, name: "A"}

	if s1 == s2 {
		fmt.Println("s1 = s2")
	} else {
		fmt.Println("s1 != s2")
	}

	type struct2 struct {
		id   int
		name string
		info map[int]int
	}

	// s3 := struct2{
	// 	id:   1,
	// 	name: "A",
	// 	info: map[int]int{
	// 		0: 1,
	// 	},
	// }

	// s4 := struct2{
	// 	id:   1,
	// 	name: "A",
	// 	info: map[int]int{
	// 		0: 1,
	// 	},
	// }

	// if s3 == s4 {
	// 	fmt.Println("s3 == s4")
	// } else {
	// 	fmt.Println("s3 != s4")
	// }

	// zero value
	var s5 struct2
	fmt.Println(s5)

}
