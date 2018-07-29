## An Introduction to Testing in Go

#### Introduction
In this tutorial we are going to look at how you can develop and run tests for your go code using the `go test` command.

---
### Go Test Files
If you have seen any go projects before, you may have noticed that most, if not all files within the project, feature a `FILE_test.go` counterpart within the same directory.

This is not an accident. These are the files which contain all of the unit tests for the project and they test all of the code within their counterparts.

##### Version 1
```
testing-in-go/
├─ main_test.go
└─ main.go
```

##### Version 2
```
testing-in-go/
├─ libs/
│   └─ calc.go
└─ tests/
    └─ calc_test.go
```

---
### A Simple Test File
Imagine we had a very simple go program that was made up of one file and featured a `calculate(`) function. This `calculate()` function simply takes in 1 number and adds 2 to it. Nice and simple to get us up and running:

```
package main

import (
	"fmt"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Hello World")
}
```
If we wished to test this we could create a `main_test.go` file within the same directory and write the following test:

```
package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
```
Note: [https://golang.org/pkg/testing/]

### Running Our Tests
Now that we have created our first go test, it’s time to run this and see if our code behaves the way we expect it to. We can execute our tests by running:

```
go test
```

This should then output something similar to the following:

```
testing-in-go> go test
PASS
ok      
testing-in-go      0.055s
```
---

### Table Driven Testing
Now that we are happy that one calculation works, we should look to improve confidence by adding a few extra test cases into our code. If we want to gradually build up a series of test cases that are always tested, we can leverage an `array` of tests like so:

```
func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
```
Here we declare a struct that contains both an input and the expected value. We then iterate through the list of tests with our `for _, test := range tests` call and check to see that our function will always return the expected results, regardless of input.

When we run our test suite now, we should see the same output as before:

```
testing-in-go> go test
PASS
ok      
testing-in-go      0.055s
```
---

### Verbose Test Output
Sometimes you may wish to see exactly what tests are running and how long they took. Thankfully, this is available if you use the `-v` flag when running your tests like so:

```
testing-in-go> go test -v
=== RUN   TestCalculate
--- PASS: TestCalculate (0.00s)
PASS
ok  
testing-in-go      0.055s
```
You can see that both our normal test and our table test ran and passed and took less than `0.00s` to execute.

---