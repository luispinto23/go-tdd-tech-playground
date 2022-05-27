---
marp: true
paginate: true
---

# What to expect?

Small introduction to the TDD process and mindset and a simple example using Go.

- 👨‍🏫 **TDD intro** - Simple introduction to TDD
<br>
<br>
- 🧪 **Sample of TDD in GO** - Basic usage of TDD in Go using Unit tests

---

# TDD

- Test-driven development
  - based on automated tests
<br>
- Useful to prevent  regressions
<br>
- Documentation for humans as to how the system should behave
<br>
- Much faster and more reliable feedback than manual testing
<br>
- Allows for a more modular design and architecture

---

# TDD cycle (Red 🔴, Green 🟢, Refactor 🔨)

After a requirement definition, a new unit test is created and the Red, Green, Refactor cycle starts.

The "red, green, refactor" cycle has 3 different stages:

- 🔴: A test for a function is failing

- 🟢: the test is passing

- 🔨: improve the code quality of the code that made the test pass, without breaking the test

---

# In practice in Go

1. Write a test for a function (ideally before writing the function) 🔴

- See the test failing
  - If no function exists, there will be a compilation error
    - Make the compiler pass by creating the function
  - If the function exists, the compiler must pass
    - Run the test, see it fails and check if it fails for the expected reason.
    - Check for the failure message to be accurate

2. Write enough code to make the test pass 🟢
3. Refactor 🔨
    - Improve the code without breaking the test

We can go through this cycle many times until a final version of the function is implemented.

---

# Writing tests (with standard testing library)

Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like `xxx_test.go`
  - Go uses this filename pattern to recognise source code files that contain tests.
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`

```go
  func TestAdd(t *testing.T) {...}
```

---

# Writing tests (with standard testing library)

A test function must declare the expected result and compare it to the returned result from the function being tested

#### The `want` and `got` pattern

```go
func TestAdd(t *testing.T) {

    var want float64 := 15

    got := calculator.Add(10, 5)

    if want != got {
        t.Errorf("want %f, got %f", want, got)
    }
}
```

---
- If `want` and `got` are different, something is wrong with our function and the test will fail

- If they are equal, the test will pass and the function ends

- Test functions don’t return anything

---

# Table-driven tests

A series of related checks can be implemented by looping over a slice of test cases.

This approach reduces the amount of repetitive code compared to repeating the same code for each test and makes it more straightforward to add more test cases.

---

```go
func TestDivide(t *testing.T) {
  t.Parallel()
  testCases := []struct {
    a, b float64
    want float64
    err  error
  }{
    {a: 10, b: 5, want: 2},
    {a: 120, b: 0, want: 0.0},
    {a: 9, b: 3, want: 3},
    {a: 99, b: 0, want: 0.0},
  }

  for _, tc := range testCases {
    got, err := calculator.Divide(tc.a, tc.b)

    if err != nil {
      t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
    }

    if tc.want != got {
      t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
    }
  }
}
```

---
### Output
```bash
--- FAIL: TestDivide (0.00s)
    calculator_test.go:27: Divide(120.000000, 0.000000): want no error for valid input, got division by zero not allowed
FAIL
exit status 1
FAIL    calculator      0.005s
```

---

# Subtests

We can use the `t.Run` method for creating subtests.

This test is a rewritten version of the earlier example using subtests:

---
```go
func TestDivide(t *testing.T) {
  t.Parallel()
  testCases := []struct {
    a, b float64
    want float64
    err  error
  }{
    {a: 10, b: 5, want: 2},
    {a: 120, b: 0, want: 0.0},
    {a: 9, b: 3, want: 3},
    {a: 99, b: 0, want: 0.0},
  }
 
  for _, tc := range testCases {
    tc := tc // capture range variable
    t.Run(fmt.Sprintf("Divide %f by %f", tc.a, tc.b), func(t *testing.T) {
      t.Parallel()
      got, err := calculator.Divide(tc.a, tc.b)

      if err != nil {
        t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
      }
    
      if tc.want != got {
        t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
      }
    })
  }
}
```

---
## Output
```bash
--- FAIL: TestDivide (0.00s)
    --- FAIL: TestDivide/Divide_99.000000_by_0.000000 (0.00s)
        calculator_test.go:40: Divide(99.000000, 0.000000): want no error for valid input, got division by zero not allowed
    --- FAIL: TestDivide/Divide_120.000000_by_0.000000 (0.00s)
        calculator_test.go:40: Divide(120.000000, 0.000000): want no error for valid input, got division by zero not allowed
FAIL
exit status 1
FAIL    calculator      0.004s
```

---

# Subtests

There's now a difference in output from the two implementations. The original implementation prints:

```bash
--- FAIL: TestDivide (0.00s)
    calculator_test.go:27: Divide(120.000000, 0.000000): want no error for valid input, got division by zero not allowed
FAIL
exit status 1
FAIL    calculator      0.005s
```

Even though there are two errors, execution of the test halts on the call to `Fatalf` and the second test case never runs.

---
The implementation using `t.Run` prints both:

```bash
--- FAIL: TestDivide (0.00s)
    --- FAIL: TestDivide/Divide_99.000000_by_0.000000 (0.00s)
        calculator_test.go:40: Divide(99.000000, 0.000000): want no error for valid input, got division by zero not allowed
    --- FAIL: TestDivide/Divide_120.000000_by_0.000000 (0.00s)
        calculator_test.go:40: Divide(120.000000, 0.000000): want no error for valid input, got division by zero not allowed
FAIL
exit status 1
FAIL    calculator      0.004s
```

`Fatal` and its siblings causes a subtest to be skipped but not its parent or subsequent subtests.

---
We also can use the `t.Run` method to create subtests with a more descriptive name and pick out a specific subtest to run.

```bash
go test -run TestDivide/floatdivision
```

---
```go
func TestDivide(t *testing.T) {
  testCases := []struct {
    name string
    a, b float64
    want float64
    err  error
  }{
    {name: "floatdivision", a: 10, b: 5, want: 2},
    {name: "dividebyzero", a: 120, b: 0, want: 0.0},
  }
 
  for _, tc := range testCases {
    tc := tc // capture range variable
    tf := func(t *testing.T) {
      t.Parallel()
      t.Log(tc.name
      got, err := calculator.Divide(tc.a, tc.b)

      if err != nil {
        t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
      }
    
      if tc.want != got {
        t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
      }
    
    }
    t.Run(tc.name, tf)
  }
}
```

---

# Test coverage

We can get the test coverage of our code by running `go test -cover`.

If we want to know what's covered by a specific test, we can use the `-coverprofile` flag and specify a file name for the profile output.
  
  ```bash
  go test -coverprofile=coverage.out
  ```
---
One way to view the coverage is to use the `go tool cover` command and pass the profile file as an argument and the `-html` flag.

```bash
go tool cover -html=coverage.out
```

this will open a web browser with the coverage report.

It's also possible to generate an HTML report with `go tool cover -html=coverage.out -o coverage.html`.

---

### references

<https://quii.gitbook.io/learn-go-with-tests/>

<https://www.youtube.com/watch?v=Bt1ZA82SF4o>

<https://go.dev/blog/subtests>

<https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721>

<https://www.gopherguides.com/articles/table-driven-testing-in-parallel>
