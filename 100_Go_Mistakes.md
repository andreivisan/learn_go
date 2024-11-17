# 100 Go Mistakes

## Ch 2 Code and project organization

### 2.1: Unintended variable shadowing

```go
var client *http.Client
if tracing {
    client, err := createClientWithTracing()
    if err != nil {
        return err
    }
    log.Println(client)
} else {
    client, err := createDefaultClient()
    if err != nil {
        return err
    }
    log.Println(client)
}
```

In this example, we first declare a client variable. Then, we use the short variable
declaration operator (:=) in both inner blocks to assign the result of the function call
to the inner client variables—not the outer one. As a result, the outer variable is
always nil.

**NOTE** This code compiles because the inner client variables are used in the
logging calls. If not, we would have compilation errors such as client
declared and not used.

The first option uses temporary variables in the inner blocks this way:

```go
var client *http.Client
if tracing {
    c, err := createClientWithTracing()
    if err != nil {
        return err
    }
    client = c
} else {
    // Same logic
}
```

The second option uses the assignment operator (=) in the inner blocks to directly
assign the function results to the client variable. However, this requires creating an
error variable because the assignment operator works only if a variable name has
already been declared. For example:

```go
var client *http.Client
var err error
if tracing {
    client, err = createClientWithTracing()
    if err != nil {
        return err
    }
} else {
    // Same logic
}
```

Also, with the second option, we can mutualize and implement error handling outside the if/else statements, as this example shows:

```go
if tracing {
    client, err = createClientWithTracing()
} else {
    client, err = createDefaultClient()
}
if err != nil {
    // Common error handling
}
```

### 2.2: Unnecessary nested code

Align the happy path to the left; you should quickly be able to scan down one column to
see the expected execution flow.

- Use happy path aligned to the left, avoid nesting.

### 2.3: Misusing init functions

An init function is a function used to initialize the state of an application. It takes no
arguments and returns no result (a func() function). When a package is initialized,
all the constant and variable declarations in the package are evaluated. Then, the init
functions are executed. 

If we have one package depending on another and both declare an init function, let’s say package main depends on redis, then the init() method in the redis package is called first since main depends on it, then the init function in main package is executed.

We can define multiple init functions per package. When we do, the execution
order of the init function inside the package is based on the source files’ alphabetical
order. For example, if a package contains an a.go file and a b.go file and both have an
init function, the a.go init function is executed first.

We shouldn’t rely on the ordering of init functions within a package. Indeed, it can be
dangerous as source files can be renamed, potentially impacting the execution order.

We can initialize packages (calling their init function) when we import them. We can do that by using the ```_``` operator this way:

```go
package main

import (
    “fmt”
    _ “foo”
)
```

- One mistake is opening a DB connection in the init function.
    
    - Error management in an init function is limited. Indeed, as an init function doesn’t return an error, one of the only ways to signal an error is to panic, leading the application to be stopped.
    
    - Another important downside is related to testing. If we add tests to this file, the init
function will be executed before running the test cases, which isn’t necessarily what
we want (for example, if we add unit tests on a utility function that doesn’t require this
connection to be created).
    
    - The last downside is that the example requires assigning the database connection pool to a global variable. Global variables have some severe drawbacks; for example:
        
        - Any functions can alter global variables within the package.
        
        - Unit tests can be more complicated because a function that depends on a global variable won’t be isolated anymore.

In most cases, we should favor encapsulating a variable rather than keeping it global.

 



 


