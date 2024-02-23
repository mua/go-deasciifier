### Deasciifier Go Package

This is a Go package for deasciifying Turkish text. It is a port of [Prof. Deniz Yuret](http://www.denizyuret.com/) . Also test cases and constants are ported from the Python version by [Emre Sevinc](https://github.com/emres/turkish-deasciifier).

### Usage of the package

```go
package main

import (
    "fmt"
    "github.com/mua/deasciifier"
)


func main() {
    deasciifier := deasciifier.NewDeasciifier()
    fmt.Println(deasciifier.Deasciify("dunyanin en guzel sehri istanbul"))
}
// Output: dünyanın en güzel şehri istanbul
```

There is also a command line tool for deasciifying text. You can install it with 
```bash
$ go get github.com/mua/deasciifier/cmd/deasciify
```
and use it like this:

```bash
$ deasciify -s "dunyanin en guzel sehri istanbul"
dünyanın en güzel şehri istanbul

$ echo "dunyanin en guzel sehri istanbul" > file.txt
$ deasciify -f file.txt
dünyanın en güzel şehri istanbul
```
