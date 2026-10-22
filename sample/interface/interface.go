package main

import "fmt"

type InterfaceA interface {
	Fn()
}

type InterfaceB struct{}

func (b *InterfaceB) Fn() {
	fmt.Println("Hello")
}

// InterfaceA를 구현한 값의 Fn 메서드를 실행한다.
func InitializeInterfaceA(a InterfaceA) {
	a.Fn()
}

// InterfaceB의 생성자
func NewInterfaceB() *InterfaceB {
	b := &InterfaceB{}
	InitializeInterfaceA(b)
	//구조체 b가 InterfaceA를 구현했으니 InterfaceA 타입으로 *B 전달 가능
	return b
}

func runInterfaceExample() {
	fmt.Println("[Go interface example]")
	var a InterfaceA = NewInterfaceB()
	a.Fn()
}
