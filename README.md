# GO

## Resources

1. [A Tour of Go](https://go.dev/tour/list)

2. [Official Docs](https://go.dev/doc/)

3. [Effective Go](https://go.dev/doc/effective_go#blank_import)

4. [Exercism](https://exercism.org/tracks/go)

5. [FreeCodeCamp Roadmap](https://www.freecodecamp.org/news/golang-developer-roadmap/)

## Installation

[Download and Install](https://go.dev/doc/install)

```bash
$ go version
```

## Hello, GO

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, GO")
}
```

### Running hello.go

There are 2 options to run hello.go:

1. Basic execution (and the simplest way)

```bash
go run hello.go
```

2. Creating and Running Executables

```bash
go build hello.go

./hello
```

## Packages

Every Go program is made up of packages.

Programs start running in package `main`.

### Imports

It's a good practice to group imports into a paranthesized, "factored" import statement.

```go
import (
	"fmt"
	"math"
)
```

### Exported names

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```

## Functions

A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

Thus `x int, y int` becomes `x, y int`

### Multiple results

A function can return any number of results.

The swap function returns two strings.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

### Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

### Variables

The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

### Basic types

Go's basic types are

```text
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

### Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

    - 0 for numeric types,
    - false for the boolean type, and
    - "" (the empty string) for strings.

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

### Type conversions

The expression T(v) converts the value v to the type T.

Some numeric conversions:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply:

```go
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```
