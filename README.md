---
theme: ./ui/themes/dracula.json
author: Luís Pinto
paging: Slide %d / %d
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
  - When refactoring code we must not be changing it's behaviour,
    - using tests will give us the confidence that we can reshape code without worrying about that
<br>
- Documentation for humans as to how the system should behave
<br>
- Much faster and more reliable feedback than manual testing
<br>
- Allows for a more modular design and architecture
  - TDD helps developers understand and learn the principles of modular design when writing tests for very small features. In this way, problems in the application’s architecture can be detected at an early stage of development.

---

# TDD cycle (Red 🔴, Green 🟢, Refactor 🔨)

After a requirement definition, a new unit test is created and the Red, Green, Refactor cycle starts.

The "red, green, refactor" cycle has 3 different stages:

- 🔴: A test for a function is failing
<br>
<br>
- 🟢: the test is passing
<br>
<br>
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
<br>

2. Write enough code to make the test pass 🟢
<br>
3. Refactor 🔨
    - Improve the code without breaking the test
<br>

We can go through this cycle many times until a final version of the function is implemented.

A function with a complicated behaviour is easily testable and implemented if we can break such behaviour in smaller parts and test accordingly.

---

# Writing tests (with standard testing library)

Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like `xxx_test.go`
  - Go uses this filename pattern to recognise source code files that contain tests.
<br>
- The test function must start with the word `Test`
<br>
- The test function takes one argument only `t *testing.T`
<br>
<br>

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

- If `want` and `got` are different, something is wrong with our function and the test will fail
<br>
<br>
- If they are equal, the test will pass and the function ends
<br>
<br>
- Test functions don’t return anything

---

# Table-driven tests

A series of related checks can be implemented by looping over a slice of test cases.

This approach reduces the amount of repetitive code compared to repeating the same code for each test
and makes it more straightforward to add more test cases.

```go
func TestDivide(t *testing.T) {
  t.Parallel()
  testCases := []struct {
    a, b float64
    want float64
    err  error
  }{
    {a: 10, b: 5, want: 2, err: nil},
    {a: 120, b: 0, want: 0.0, err: errors.New("division by zero not allowed")},
    {a: 99, b: 0, want: 0.0, err: errors.New("division by zero not allowed")},
    {a: 9, b: 3, want: 3, err: nil},
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

```go
func TestDivide(t *testing.T) {
  t.Parallel()
  testCases := []struct {
    a, b float64
    want float64
    err  error
  }{
    {a: 10, b: 5, want: 2, err: nil},
    {a: 120, b: 0, want: 0.0, err: errors.New("division by zero not allowed")},
    {a: 99, b: 0, want: 0.0, err: errors.New("division by zero not allowed")},
    {a: 9, b: 3, want: 3, err: nil},
  }
 
  for _, tc := range testCases {
    t.Run(fmt.Sprintf("Divide %f by %f", tc.a, tc.b), func(t *testing.T) {
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
### Output
```bash
--- FAIL: TestDivide (0.00s)
    --- FAIL: TestDivide/Divide_120.000000_by_0.000000 (0.00s)
        calculator_test.go:52: Divide(120.000000, 0.000000): want no error for valid input, got division by zero not allowed
    --- FAIL: TestDivide/Divide_99.000000_by_0.000000 (0.00s)
        calculator_test.go:52: Divide(99.000000, 0.000000): want no error for valid input, got division by zero not allowed
FAIL
exit status 1
FAIL    calculator      0.005s
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

Even though there are two errors, execution of the test halts on the call to `Fatalf` and the second test never runs.


The implementation using `t.Run` prints both:

```bash
--- FAIL: TestDivide (0.00s)
    --- FAIL: TestDivide/Divide_120.000000_by_0.000000 (0.00s)
        calculator_test.go:52: Divide(120.000000, 0.000000): want no error for valid input, got division by zero not allowed
    --- FAIL: TestDivide/Divide_99.000000_by_0.000000 (0.00s)
        calculator_test.go:52: Divide(99.000000, 0.000000): want no error for valid input, got division by zero not allowed
FAIL
exit status 1
FAIL    calculator      0.005s
```

`Fatal` and its siblings causes a subtest to be skipped but not its parent or subsequent subtests.

---

### references

<https://go.dev/blog/subtests>>

<https://quii.gitbook.io/learn-go-with-tests/>

<https://www.youtube.com/watch?v=Bt1ZA82SF4o>
