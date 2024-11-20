package rows

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	row := New(3, "*")

	if len(row) != 3 {
		t.Errorf("expected length 3")
	}

	for i := 0; i < 3; i++ {
		if row[i] != "*" {
			t.Errorf("expected row[%v] == *, got %v", i, row[i])
		}
	}
}

func ExampleRow_String() {
	row1 := New(3, "*")
	row2 := New(4, "#")
	row3 := New(1, "A")

	fmt.Println(row1)
	fmt.Println(row2)
	fmt.Println(row3)

	// Output:
	// | * | * | * |
	// | # | # | # | # |
	// | A |
}

func ExampleRow_ContainsOnly_positive() {
	row1 := New(3, "*")
	row2 := New(4, "#")
	row3 := New(1, "A")

	fmt.Println(row1.ContainsOnly("*"))
	fmt.Println(row2.ContainsOnly("#"))
	fmt.Println(row3.ContainsOnly("A"))

	// Output:
	// true
	// true
	// true
}

func ExampleRow_ContainsOnly_negative() {
	row1 := Row{"*", "#", "*"}
	row2 := Row{"#", "#", "A", "#"}
	row3 := New(3, "*")

	fmt.Println(row1.ContainsOnly("*"))
	fmt.Println(row2.ContainsOnly("#"))
	fmt.Println(row3.ContainsOnly("A"))

	// Output:
	// false
	// false
	// false
}
