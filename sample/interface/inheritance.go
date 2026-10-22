package main

import "fmt"

type InheritanceA interface {
	Fn()
}

// 부모 클래스 역할
type Base struct{}

// 순수 가상 함수의 오류를 흉내 낸 구현
func (a *Base) Fn() {
	panic("pure virtual function called")
}

// InheritanceA의 생성자 역할
func NewBase() *Base {
	a := &Base{}
	a.Fn()
	return a
}

// 자식 클래스 역할
type InheritanceB struct {
	*Base
}

func (b *InheritanceB) Fn() {
	fmt.Println("Hello")
}

// InheritanceB의 생성자 역할
func NewInheritanceB() *InheritanceB {
	b := &InheritanceB{
		Base: NewBase(),
	}
	return b
}

func runInheritanceExample() {
	fmt.Println("[C++ inheritance example]")

	// 이 예제에서 발생하는 panic 이후에도 인터페이스 예제를 실행한다.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()

	var a InheritanceA = NewInheritanceB()
	a.Fn()
}
