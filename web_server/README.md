# HTTP Server in GO

1. When using a standard fileserver, the path to a file on disk is the same as its URL path. An exception is that index.html is served from / instead of /index.html.

For example:

- /index.html will be served from /
- /pages/index.html will be served from /pages
- /pages/about/index.html will be served from /pages/about