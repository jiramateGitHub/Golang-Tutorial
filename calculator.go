package library

// Calculator is structure for calculator
type Calculator struct {
}

// add
func (c *Calculator) Add(num1 int, num2 int) int {
	return num1 + num2
}

// subtract
func (c *Calculator) Subtract(num1 int, num2 int) int {
	return num1 - num2
}

// multiply
func (c *Calculator) Multiply(num1 int, num2 int) int {
	return num1 * num2
}

// divide
func (c *Calculator) Divide(num1 int, num2 int) int {
	if num2 == 0 {
		return 0
	}

	return num1 / num2
}
