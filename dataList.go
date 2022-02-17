package main

import (
	"fmt"
)

type ListData interface {
	Length() int
	nameOfData() string
}

func fibonachi(listData ListData) {
	var a = 0
	var b = 1
	var c = a + b

	fmt.Printf("%d %d ", a, b)
	for i := 2; i < listData.Length(); i++ {
		fmt.Printf("%d ", c)
		a = b
		b = c
		c = a + b
	}
	fmt.Println("- ", listData.nameOfData())
}

func doubleData(listData ListData) {
	var a = 0
	var b = 0
	for i := 0; i < listData.Length(); i++ {
		fmt.Printf("%d ", a)
		b = a + 2
		a = b
	}
	fmt.Println("- ", listData.nameOfData())
}

type Data struct {
	DataLength int
	Name       string
}

func (data Data) Length() int {
	return data.DataLength
}

func (data Data) nameOfData() string {
	return data.Name
}

func main() {
	var Fibonachi Data
	Fibonachi.DataLength = 12
	Fibonachi.Name = "Fibonaci Data"

	var DoubleData Data
	DoubleData.DataLength = 12
	DoubleData.Name = "Double Data"

	fibonachi(Fibonachi)
	doubleData(DoubleData)

}
