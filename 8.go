package main

import (
	"fmt"
	"sync"
)

type Boiler struct {
	empty  bool
	boiled bool
}

var instance *Boiler
var once sync.Once

func GetInstance() *Boiler {
	once.Do(func() {
		instance = &Boiler{}
	})
	return instance
}

func (c *Boiler) IsEmpty() bool {
	return c.empty
}

func (c *Boiler) IsBoiled() bool {
	return c.boiled
}

func (c *Boiler) Fill() {
	if c.empty {
		c.empty = false
		fmt.Println("容器装满了")
	}
}

func (c *Boiler) Drain() {
	if c.empty == false && c.boiled {
		c.empty = true
		c.boiled = false
		fmt.Println("倒入模具了")
	}
}

func (c *Boiler) Boil() {
	if c.empty == false && c.boiled == false {
		fmt.Println("巧克力煮开了")
		c.boiled = true
	}
}

func main() {
	instance = GetInstance()
	fmt.Println(instance.IsBoiled(), instance.IsEmpty())
	instance.Fill()
	instance.Boil()
	// 在倒入模具之前我们再新建一个实例，看看会不会把之前操作全都抛弃。
	instance = GetInstance()
	// 看看是否保持了煮非空的状态
	fmt.Println(instance.IsBoiled(), instance.IsEmpty())
	instance.Drain()
	fmt.Println(instance.IsBoiled(), instance.IsEmpty())
}
