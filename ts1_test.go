package lab1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask1(t *testing.T) {
	res, err := Task1("42 21 - 322 * 56 +")
	if assert.Nil(t, err) {
		assert.Equal(t,"+56 * 322 - 42 21", res)
	}
	fmt.Println(res)
}

func ExampleTask1() {
	res, _ := Task1("2 2+")
	fmt.Println(res)

	// Output:
	// +2 2
}
