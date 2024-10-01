package design

import (
	"fmt"
	"testing"
)

// 模板模式
type Cooker interface {
	fire()
	cooke()
	outfire()
}

type CookMenu struct{}

func (CookMenu) fire() {
	fmt.Println("点火")
}

func (CookMenu) cooke() {}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

type ToMaTo struct {
	CookMenu
}

func (t *ToMaTo) cooke() {
	fmt.Println("cook tomato")
}

type Aggs struct {
	CookMenu
}

func (a *Aggs) cooke() {
	fmt.Println("do aggs")
}

func TestTemplate(t *testing.T) {
	// 做西红柿
	xihongshi := &ToMaTo{}
	doCook(xihongshi)

	fmt.Println("\n=====> 做另外一道菜")
	// 做炒鸡蛋
	aggs := &Aggs{}
	doCook(aggs)

}
