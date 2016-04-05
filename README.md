Gobox
=====

A (**work in progress**) lightweight, no-frills file sharing server written in Go. Uses [Gorilla Mux](http://github.com/gorilla/mux) for routing.

Usage
=====

(Only tested on Linux/OSX).

1. Install Gorilla Mux with `go get github.com/gorilla/mux`
2. Build with `go build`.
3. Copy `gobox` executable to the folder you want to share.
4. Run `./gobox`.

Parameters:

| param | type | description |
| ----- | ---- | ----------- |
| `--address` | string | listen address (default: 127.0.0.0) |
| `--port` | int | listen port (default: 3000) |
| `--verbose` | bool | log requests | 

TODO
====

- Add fancy schmancy design to make it not look like an 80's page
- Add file upload function

License
=======

The MIT License (MIT)
Copyright (c) 2016 KixPanganiban

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.