package main

import "fmt"

type Product struct {
	ground string
	cement string
	roof   string
}

func (p *Product) Cement() string {
	return p.cement
}

func (p *Product) SetCement(cement string) {
	p.cement = cement
}

func (p *Product) Roof() string {
	return p.roof
}

func (p *Product) SetRoof(roof string) {
	p.roof = roof
}

func (p *Product) Ground() string {
	return p.ground
}

func (p *Product) SetGround(ground string) {
	p.ground = ground
}

type Builder interface {
	BuildGround()
	BuildCement()
	BuildRoof()
	BuilderProduct() *Product
}

type ConcreteBuilder struct {
	p *Product
}

func (c *ConcreteBuilder) BuildGround() {
	c.p.SetGround("建造地基")
	fmt.Println(c.p.Ground())
}
func (c *ConcreteBuilder) BuildCement() {
	c.p.SetCement("建造房子")
	fmt.Println(c.p.Cement())
}
func (c *ConcreteBuilder) BuildRoof() {
	c.p.SetRoof("建造屋顶")
	fmt.Println(c.p.Roof())
}

func (c *ConcreteBuilder) BuilderProduct() *Product {
	return c.p
}

type Director struct {
	builder Builder
}

func (d *Director) Construst() Product {
	d.builder.BuildGround()
	d.builder.BuildCement()
	d.builder.BuildRoof()
	return *d.builder.BuilderProduct()
}

func main21() {
	builder := &ConcreteBuilder{p: &Product{}}
	direactor := &Director{builder: builder}
	direactor.Construst()
}
