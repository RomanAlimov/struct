package main

import "fmt"

type User struct { // структуры с типами
	name string
	age  int
}

func (u *User) gname() bool { // функция с указателем на структуры
	a := u.age  //указатель на возраст из структуры
	tr := false //флаг

	if a < 18 {
		tr = false
	} else if a >= 18 {
		tr = true
	}
	return tr
}

func main() {
	user := User{name: "Roman", age: 17} // передача данных из стрк
	fmt.Print(user)
	fmt.Print(user.gname()) //функция
}
