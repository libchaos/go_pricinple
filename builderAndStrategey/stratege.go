package main

import "fmt"

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

type Duck struct {
	fly FlyBehavior
	quack QuackBehavior
}

func (d *Duck) Swim()  {
	fmt.Println("鸭子游泳")
}

func (d *Duck) Display(behavior FlyBehavior, quackBehavior QuackBehavior)  {
	behavior.Fly()
	quackBehavior.Quack()
}

type FlyWithWings struct{}

func (f *FlyWithWings) Fly()  {
	fmt.Println("鸭子用翅膀飞")
}

type FlyNoWay struct{}

func (f *FlyNoWay) Fly()  {
	fmt.Println("鸭子嘎嘎叫")
}

// 具体鸭子叫
type Quack struct {}

func (f *Quack) Quack ()  {
    fmt.Println("鸭子嘎嘎叫")
}
// 具体鸭子叫
type Squeak struct {}

func (f *Squeak) Quack ()  {
    fmt.Println("鸭子咔咔叫")
}
// 具体鸭子叫
type Mute struct {}

func (f *Mute) Quack ()  {
    fmt.Println("鸭子不能叫")
}
// 家鸭
type ReadHead struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (r *ReadHead) Display ()  {
    r.Swim()
    r.Duck.Display(r.fly, r.quack)
}
// 木头鸭子
type Wooden struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (r *Wooden) Display ()  {
    r.Swim()
    r.Duck.Display(r.fly,r.quack)
}
// 野鸭
type Mallard struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (m *Mallard) Display ()  {
    m.Swim()
    m.Duck.Display(m.fly, m.quack)
}
// 橡胶鸭子
type Rubber struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (r *Rubber) Display ()  {
    r.Swim()
    r.Duck.Display(r.fly, r.quack)
}


func main(){
	// 新建鸭子的各种行为
	flynoway := &FlyNoWay{}
	flayWihtwings := &FlyWithWings{}
	quack := &Quack{}
	sqeak := &Squeak{}
	mute := &Mute{}
	// 对于以下四种鸭子，我们按需组合各种技能，是不是特别方便
	duck := ReadHead{
			Duck:  Duck{},
			fly:   flayWihtwings,
			quack: quack,
	}
	duck.Display()
	mallard := Mallard {
			Duck:  Duck{},
			fly:   flayWihtwings,
			quack: quack,
	}
	mallard.Display()
	rub := Rubber {
			Duck:  Duck{},
			fly:   flynoway,
			quack: sqeak,
	}
	rub.Display()
	wooden := Wooden{
			Duck:  Duck{},
			fly:   flynoway,
			quack: mute,
	}
	wooden.Display()
}
