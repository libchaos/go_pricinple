package main

import "fmt"

// 最抽象的一个工厂接口
type Factory interface {
	NewTV() Television
	NewRefrigerator() Refrigerator
}

// 两个工厂都有的产品的接口
type Television interface {
	DoSomething()
}

type Refrigerator interface {
	DoSomething()
}

type TCLTV struct {
}

func (TCLTV) DoSomething() {
	fmt.Println("TCL电视在Do Something")
}

type TCLRef struct {
}

func (TCLRef) DoSomething() {
	fmt.Println("TCL空调在do something")
}

type TCLFactory struct {
}

func (TCLFactory) NewTV() Television {
	return TCLTV{}
}

func (TCLFactory) NewRefrigerator() Refrigerator {
	return TCLRef{}
}

type MediaTV struct {
}

func (MediaTV) DoSomething() {
	fmt.Println("美的电视在do something")
}

type MediaRef struct{}

func (MediaRef) DoSomething() {
	fmt.Println("美的空调在do something")
}

type MediaFactory struct {
}

func (MediaFactory) NewTV() Television {
	return MediaTV{}
}

func (MediaFactory) NewRefrigerator() Refrigerator {
	return MediaRef{}
}

func main() {
	var (
		factory Factory
	)
	// 这里不管是TCL工厂还是美的工厂，因为他们都实现了Factory的接口，
	// 所以这两个类都可以直接当做Factory对象来直接使用。
	factory = &TCLFactory{}
	ref := factory.NewRefrigerator()
	ref.DoSomething()
	tv := factory.NewTV()
	tv.DoSomething()

	factory = MediaFactory{}
	ref = factory.NewRefrigerator()
	ref.DoSomething()
	tv = factory.NewTV()
	tv.DoSomething()
}
