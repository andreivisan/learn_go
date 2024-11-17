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


