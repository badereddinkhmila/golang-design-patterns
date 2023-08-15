package creational

import "fmt"

/*********************************************************************/
/*************************** Interface *******************************/
/*********************************************************************/

type IClonable interface {
	print(string)
	clone() IClonable
}

/*********************************************************************/
/*************************** Implementation **************************/
/*********************************************************************/

type Student struct {
	Firstname string
	Age       int
}

func (s *Student) clone() IClonable {
	return &Student{
		Firstname: s.Firstname,
		Age:       s.Age,
	}
}

func (s *Student) print(indentation string) {
	fmt.Println(indentation + s.Firstname + "_clone")
}

type Classroom struct {
	Floor    string
	Tag      string
	students []IClonable
}

func (c *Classroom) clone() IClonable {
	class := &Classroom{Floor: c.Floor + "_clone", Tag: c.Tag + "_clone"}
	var students []IClonable
	for _, v := range c.students {
		copy := v.clone()
		students = append(students, copy)
	}
	class.students = students
	return class
}

func (c *Classroom) print(indentation string) {
	fmt.Println(indentation + c.Tag)
	for _, i := range c.students {
		i.print(indentation + "[Child] ")
	}
}
