package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	temp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = temp

	expected := "Title: Cat in the Hat, Author: Dr.Seuss, Pages: 20, Read: false, Owner: {name:Megan age:39}\n"
	if string(out) != expected {
		t.Errorf("\nExpected: %s\ngot     :  %s", expected, out)
	} else {
		fmt.Println("Passed main output test")
	}

}

func TestBookString(t *testing.T) {
	tests := [5]struct {
		def    string
		book   Book
		output string
	}{
		{
			"normal input",
			Book{"Lord of the Rings", "J.R.R. Tolkien", 355, true, &Person{"Riley", 25}},
			"Title: Lord of the Rings, Author: J.R.R. Tolkien, Pages: 355, Read: true, Owner: {name:Riley age:25}",
		},
		{
			"missing Person field",
			Book{"Chronicles of Narnia", "C.S. Lewis", 388, false, &Person{name: "Indy"}},
			"Title: Chronicles of Narnia, Author: C.S. Lewis, Pages: 388, Read: false, Owner: {name:Indy age:0}",
		},
		{
			"empty Person struct",
			Book{"Lord of the Rings", "J.R.R. Tolkien", 355, true, &Person{}},
			"Title: Lord of the Rings, Author: J.R.R. Tolkien, Pages: 355, Read: true, Owner: {name: age:0}",
		},
		{
			"missing Person struct",
			Book{title: "Chronicles of Narnia", author: "C.S. Lewis", pages: 388, read: false},
			"Title: Chronicles of Narnia, Author: C.S. Lewis, Pages: 388, Read: false, Owner: nil",
		},
		{
			"Missing book fields",
			Book{title: "Lord of the Rings", owner: &Person{"Riley", 25}},
			"Title: Lord of the Rings, Author: , Pages: 0, Read: false, Owner: {name:Riley age:25}",
		},
		// Add more test cases here if needed
	}

	for i, test := range tests {
		t.Run(test.output, func(t *testing.T) {
			actualOutput := fmt.Sprintf("%v", test.book)
			if actualOutput != test.output {
				t.Errorf("Expected: %v, but got: %v", test.output, actualOutput)
			} else {
				fmt.Printf("Pass Test %v: %s\n", i+1, test.def)
			}
		})
	}
}
