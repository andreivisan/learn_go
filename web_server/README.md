# HTTP Server in GO

1. When using a standard fileserver, the path to a file on disk is the same as its URL path. An exception is that index.html is served from / instead of /index.html.

For example:

- /index.html will be served from /
- /pages/index.html will be served from /pages
- /pages/about/index.html will be served from /pages/about

2. Handler Review

2.1. Handler

An http.Handler is any defined type that implements the set of methods defined by the Handler interface, specifically the ServeHTTP method.

The ServeMux you used in the previous exercise is an http.Handler.

You will typically use a Handler for more complex use cases, such as when you want to implement a custom router, middleware, or other custom logic.

2.2. HandlerFunc

HandlerFunc is a function wrapper so that when we call HandlerFunc(f) then the ServeHTTP method of Handler will call f(w, r).

If we didn't have HandlerFunc I would have to create a type that implements the Handler interface.

```go
type MyHandler struct {
    header      string
    otherParams string
    // Add any other fields you need
}

// Implement the ServeHTTP method to satisfy the http.Handler interface
func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Set the Content-Type header
    w.Header().Set("Content-Type", handler.header)

    // Write the status code
    w.WriteHeader(http.StatusOK)

    // Write the response body
    body := []byte("Hello from MyHandler!")
    w.Write(body)

    // You can access otherParams or other fields as needed
    // For example, use handler.otherParams for custom logic
}
```


You'll typically use a HandlerFunc when you want to implement a simple handler. The HandlerFunc type is just a function that matches the ServeHTTP signature above.

