---
GO TDD Primer
- TDD intro
- Basic usage of TDD in Go

---

# TDD

- Test-driven development
  - automated tests
- Prevent  regressions
  - When refactoring code you must not be changing behaviour
  - Confidence that you can reshape code without worrying about changing behaviour
- Documentation for humans as to how the system should behave
- Much faster and more reliable feedback than manual testing
- Allows for a more modular design and architecture

## TDD cycle (Red, Green, Refactor)

The "red, green, refactor" cycle has 3 different stages:
    - Red: a test for a function is failing
    - Green: the test is passing
    - Refactor: improve the code quality of the code that made the test pass, without breaking the test

---

### in Practice

- Write a test for a function (ideally before writing the function)
  - See the test failing
    - Compilation error if no function exists
      - Make the compiler pass and the test fail
    - Compile and fail if there's already a function
      - Run the test, see it fails and check if it fails for the expected reason.
        - Check for the failure message to be accurate
- Write enough code to make the test pass
- Refactor
  - Improve the code without breaking the test

We can go through this cycle many times until a final version of the function is implemented.

A function with a complicated behavior is easily testable and implemented if we can break such behavior in smaller parts and test accordingly.

---

## Writing tests (with standard testing library)

Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like `xxx_test.go`
  - Go uses this filename pattern to recognize source code files that contain tests.

- The test function must start with the word `Test`

- The test function takes one argument only `t *testing.T`

    ```go
    func TestAdd(t *testing.T) {...}
    ```

- A test function must declare the expected result and compare it to the returned result from the function being tested
  - The `want` and `got` pattern

```go
func TestAdd(t *testing.T) {

    var want float64 = 15

    got := calculator.Add(10, 5)

    if want != got {
        t.Errorf("want %f, got %f", want, got)
    }
}
```

- If want and got are different, something is wrong with our function and the test will fail
- If they are the equal, the test will pass and the function ends

- Test functions don’t return anything

## TODO: *t Testing definição

---

### Table-driven tests

A series of related checks can be implemented by looping over a slice of test cases:

```go
func TestAdd(t *testing.T) {
    t.Parallel()
    type testCase struct {
        a, b float64
        want float64
    }
    
    testCases := []testCase{
        {a: 2, b: 2, want: 4},
        {a: 1, b: 1, want: 2},
        {a: 5, b: 0, want: 5},
    }

    for _, tc := range testCases {
        got, err := calculator.Add(tc.a, tc.b)

        if err != nil {
            t.Fatalf("Add(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
        }
        if tc.want != got {
            t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
        }
    }
}
```

reduces the amount of repetitive code compared to repeating the same code for each test and makes it straightforward to add more test cases.

---

### subtests

 `t.Run` method for creating subtests. This test is a rewritten version of our earlier example using subtests:

```go
func TestTime(t *testing.T) {
    testCases := []struct {
        gmt  string
        loc  string
        want string
    }{
        {"12:31", "Europe/Zuri", "13:31"},
        {"12:31", "America/New_York", "7:31"},
        {"08:08", "Australia/Sydney", "18:08"},
    }
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
            loc, err := time.LoadLocation(tc.loc)
            if err != nil {
                t.Fatal("could not load location")
            }
            gmt, _ := time.Parse("15:04", tc.gmt)
            if got := gmt.In(loc).Format("15:04"); got != tc.want {
                t.Errorf("got %s; want %s", got, tc.want)
            }
        })
    }
}
```

The first thing to note is the difference in output from the two implementations. The original implementation prints:

```bash
--- FAIL: TestTime (0.00s)
    time_test.go:62: could not load location "Europe/Zuri"
```

Even though there are two errors, execution of the test halts on the call to `Fatalf` and the second test never runs.

The implementation using `Run` prints both:

```bash
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
        time_test.go:84: could not load location
    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
        time_test.go:88: got 07:31; want 7:31
```

`Fatal` and its siblings causes a subtest to be skipped but not its parent or subsequent subtests.

Another thing to note is the shorter error messages in the new implementation. Since the subtest name uniquely identifies the subtest there is no need to identify the test again within the error messages.

---

### avoid different tests for the same happy path

- ex: two test cases of adding positive numbers
  - this is unnecessary because the base condition doesn't change

### references

<https://go.dev/blog/subtests>

<https://quii.gitbook.io/learn-go-with-tests/>

<https://www.youtube.com/watch?v=Bt1ZA82SF4o>
