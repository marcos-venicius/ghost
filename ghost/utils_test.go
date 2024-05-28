package ghost

import (
	"testing"
)

func TestFilter(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}

	result := filter(array, func(n int) bool { return n%2 == 0 })

	if len(array) != 5 {
		t.Fatalf(`Base array should not change your size: Expected 6 received %d`, len(array))
	}

	if len(result) != 2 {
		t.Fatal("Result should have only 2 numbers")
	}
}

func TestSplitString(t *testing.T) {
	url := "/users/:id/info"

	route := splitUrl(url)

	if len(route) != 3 {
		t.Fatalf(`Expected 3 received %d`, len(route))
	}

	if route[0] != "users" {
		t.Fatalf(`Expected users received %v`, route[0])
	}

	if route[1] != ":id" {
		t.Fatalf(`Expected :id received %v`, route[1])
	}

	if route[2] != "info" {
		t.Fatalf(`Expected info received %v`, route[2])
	}
}
