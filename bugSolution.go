```go
package main

import "fmt"

func main() {
    var m map[string]int
    if m == nil {
        m = make(map[string]int)
    }
    fmt.Println(m["key"]) //Safe access, will output 0 if key is not present
}
```