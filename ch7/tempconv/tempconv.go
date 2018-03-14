package tempconv

import (
	"learn/ch2/tempconv0"
	"fmt"
)

type celsiusFlag struct {
	tempconv0.Celsius
}

func (f *celsiusFlag) Set(s string) error  {
	var value float64
	var unit string
	fmt.Sscan(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = tempconv0.Celsius(value)
		return nil
	case "F":
		f.Celsius = tempconv0.FToC(tempconv0.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}