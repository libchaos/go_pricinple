package main

import "fmt"

type ICar interface {
	GetName() string
	GetPrice() int
}

type BenzCar struct {
	name  string
	price int
}

func (b BenzCar) GetName() string {
	return b.name
}

func (b BenzCar) GetPrice() int {
	return b.price
}

type FinaceBenzCar struct {
	BenzCar
}

func (b FinaceBenzCar) GetPrice() int {
	selfPrice := b.price
	var finace int
	if selfPrice >= 100 {
		finace = selfPrice + selfPrice*5/100

	} else if selfPrice >= 50 {
		finace = selfPrice + selfPrice*2/100
	} else {
		finace = selfPrice
	}
	return finace
}

func main1() {
	var (
		list []ICar
	)
	list = []ICar{}
	list = append(list, &BenzCar{"AMG", 343})
	list = append(list, &BenzCar{"V", 60})
	for _, v := range list {
		fmt.Println("car name is ", v.GetName(), "\tprice is ", v.GetPrice())
	}
}
