# Go言語用和暦ライブラリ

## 使い方

```go
package main

import (
    "time"
    "fmt"
    "wareki"
)

func main() {
    // Go言語バージョン1のリリース日
    firstReleaseDateOfMajorVersionOfGolang := time.Date(2012, 3, 28, 0, 0, 0, 0, time.UTC)

    // 和暦を取得
    wareki := wareki.NewWareki(firstReleaseDateOfMajorVersionOfGolang)

    // 和暦の文字列を取得
    fmt.Println(wareki.String()) // 平成24年3月28日
    fmt.Printf("%s\n", wareki) // 平成24年3月28日

    // 和暦から西暦の日付を取得
    d := wareki.Date // time.Time型
    fmt.Printf("%v\n", d) //
}
```