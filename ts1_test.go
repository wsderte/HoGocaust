package lab1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask1(t *testing.T) {
	res, err := Task1("4 2-3*5+")
	if assert.Nil(t, err) {
		assert.Equal(t,"+5 *3 -4 2", res)
	}
	fmt.Println(res)
}

func ExampleTask1() {
	res, _ := Task1("2 2+")
	fmt.Println(res)

	// Output:
	// +2 2
}
