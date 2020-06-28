package main

import "fmt"

type Animal interface {
	dance()
}

type Rabbit struct {
}

func (r Rabbit) dance() {
	fmt.Println("兔子跳舞")
}

type Dog struct {
}

func (d Dog) dance() {
	fmt.Println("狗跳舞")
}

type Lion struct{}

func (l Lion) dance() {
	fmt.Println("狮子跳舞")
}

type Person struct {
	ani Animal
}

func (p Person) WalkAnimal() {
	fmt.Println("人开始遛动物")
	p.ani.dance()
}

func main2() {
	person := Person{ani: &Dog{}}
	person.WalkAnimal()
}
