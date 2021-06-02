//greeting.go
package library

import "time"

// Calculator is structure for calculator
type Greeting struct {
}

func (c *Greeting) GetGreetingByHours(time time.Time) string {
	hours := time.Hour()
	if hours < 12 {
		return "Good morning"
	}
	if hours <= 18 {
		return "Good afternoon"
	}
	return "Good evening"
}
