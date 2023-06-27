package main

import (
	"testing"
)

type Demo int
type Demo1 string

const (
	A Demo = iota
	B
	C
	D Demo1 = "hello"
	E
	F Demo = iota

	// TEST

	G
	H Demo = 3
	I
)

func TestEnum(t *testing.T) {
	println("A:", A)
	println("B:", B)
	println("C:", C)
	println("D:", D)
	println("E:", E)
	println("F:", F)
	println("G:", G)
	println("H:", H)
	println("I:", I)
}
