package callogic

type calculator struct {
	a float64
	b float64
}

func Newvalues(a, b float64) *calculator {
	return &calculator{a, b}
}

func (c *calculator) Add() float64 {
	return c.a + c.b
}

func (c *calculator) Subtract() float64 {
	return c.a - c.b
}
func (c *calculator) Multiply() float64 {
	return c.a * c.b
}
func (c *calculator) Divide() float64 {
	if c.b == 0 {
		return 0
	}
	return c.a / c.b
}
