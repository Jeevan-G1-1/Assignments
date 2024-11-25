package main

import (
	"calculator/calLogic"
	"calculator/memory"
	"fmt"
)

func main() {
	v := callogic.Newvalues(3, 4)
	fmt.Println(v.Add())
	fmt.Println(v.Subtract())
	m := memory.Factory()
	m.Setmemory(0)
	fmt.Println(m.AddtoMemory(5))
	fmt.Println(m.SubfromMemory(3))
	m.ResetMemory()
	fmt.Println(m.AddtoMemory(0))

}
