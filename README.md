# strdice
Go implementation of [Dice coefficient](https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient) to find similarity between two strings.

## Install
```bash
go get github.com/FR0ST1N/strdice
```

## Example
```go
package main

import (
        "fmt"
        "github.com/FR0ST1N/strdice"
)

func main() {
        fmt.Print(strdice.Compute("mouse", "mouse trap")) // 0.66
}
```

## License
- [MIT](LICENSE)
