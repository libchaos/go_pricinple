package main

import "fmt"

type girl interface {
	weight()
}

type FatGirl struct{}
type ThinGirl struct{}

func (f FatGirl) weight() {
	fmt.Println("fat girl")
}

func (t ThinGirl) weight() {
	fmt.Println("thin girl")
}

type GirlFactory struct{}

func (*GirlFactory) CreateGirl(like string) girl {
	switch like {
	case "fat":
		return &FatGirl{}
	case "thin":
		return &ThinGirl{}
	default:
		return nil
	}
}

func main22() {
	cg := &GirlFactory{}
	t := cg.CreateGirl("fat")
	t.weight()
}
