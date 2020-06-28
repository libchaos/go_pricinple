package main

import "fmt"

type Customer struct {
}

// func (c Customer) Shopping(shop DD) {
// 	fmt.Println(shop.sell())
// }

// func (c Customer) Shopping(shop AM) {
// 	fmt.Println(shop.sell())
// }

type Shop interface {
	sell() string
}

func (c Customer) Shopping(shop Shop) {
	fmt.Println(shop.sell())
}
