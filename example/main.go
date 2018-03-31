package main

import (
    "time"
    "fmt"
    "github.com/hideshi/go-wareki"
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
    fmt.Printf("%v\n", d) // 2012-03-28 00:00:00 +0000 UTC
}
