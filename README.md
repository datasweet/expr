
# Expr
[![Circle CI](https://circleci.com/gh/datasweet/expr.svg?style=svg)](https://circleci.com/gh/datasweet/datatable) [![GoDoc](https://godoc.org/github.com/datasweet/expr?status.png)](https://godoc.org/github.com/datasweet/expr) [![GitHub stars](https://img.shields.io/github/stars/datasweet/expr.svg)](https://github.com/datasweet/expr/stargazers)
[![GitHub license](https://img.shields.io/github/license/datasweet/expr.svg)](https://github.com/datasweet/expr/blob/master/LICENSE)

[![datasweet-logo](https://www.datasweet.fr/wp-content/uploads/2019/02/datasweet-black.png)](http://www.datasweet.fr)

Expr is an engine that can evaluate expressions. 
Expr is a fork from https://github.com/antonmedv/expr

We used the initial package to add custom formula in our package https://github.com/datasweet/datatable. 
We forked the initial project to fit our needs: operators and functions must process scalar or slices.

```go
column + 1
```
with column = [1,2,3,4,5]
output = [2,3,4,5,6]

Also, we removed the map nodes (ie {"foo": "bar}), the struct evaluation, and the type checking before evaluation.

## Installation
```go
go get github.com/datasweet/expr
```

## Who are we ?
We are Datasweet, a french startup providing full service (big) data solutions.

## Questions ? problems ? suggestions ?
If you find a bug or want to request a feature, please create a [GitHub Issue](https://github.com/datasweet/expr/issues/new).

## License
```
MIT License

Copyright (c) 2018 Anton Medvedev
Portions Copyright (c) 2019 Datasweet

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

