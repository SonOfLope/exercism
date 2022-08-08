package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	v, ok := fnb.(FancyNumber)
	if ok {
		str, _ := strconv.Atoi(v.Value())
		return str
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %d.0", ExtractFancyNumber(fnb))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	int, ok := i.(int)
	if ok {
		return DescribeNumber(float64(int))
	}
	f, ok := i.(float64)
	if ok {
		return DescribeNumber(f)
	}
	nb, ok := i.(NumberBox)
	if ok {
		return DescribeNumberBox(nb)
	}
	fancy, ok := i.(FancyNumberBox)
	if ok {
		return DescribeFancyNumberBox(fancy)
	}
	return "Return to sender"
}
