# GO

## Resources

1. [A Tour of Go](https://go.dev/tour/list)

2. [Official Docs](https://go.dev/doc/)

3. [Exercism](https://exercism.org/tracks/go)

4. [FreeCodeCamp Roadmap](https://www.freecodecamp.org/news/golang-developer-roadmap/)


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

Programs start running in package ```main```.

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

Thus ```x int, y int``` becomes ```x, y int```

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


