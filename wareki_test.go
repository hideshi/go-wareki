package wareki

import (
    "testing"
    "time"
    "fmt"
)

func TestGetWareki01_01(t *testing.T) {
    d := time.Date(1868, 1, 25, 0, 0, 0, 0, time.UTC)
    g, y, _ := GetWareki(d)
    if g != "明治" {
        t.Errorf("%s is not 明治", g)
    }
    if y != 1 {
        t.Errorf("%d is not valid Wareki year", y)
    }
}

func TestGetWareki01_02(t *testing.T) {
    d := time.Date(1912, 7, 29, 23, 59, 59, 999999999, time.UTC)
    g, y, _ := GetWareki(d)
    if g != "明治" {
        t.Errorf("%s is not 明治", g)
    }
    if y != 45 {
        t.Errorf("%d is not valid Wareki year", y)
    }
}

func TestGetWareki02_01(t *testing.T) {
    d := time.Date(1912, 7, 30, 0, 0, 0, 0, time.UTC)
    g, y, _ := GetWareki(d)
    if g != "大正" {
        t.Errorf("%s is not 大正", g)
    }
    if y != 1 {
        t.Errorf("%d is not valid Wareki year", y)
    }
}

func TestGetWareki02_02(t *testing.T) {
    d := time.Date(1926, 12, 24, 23, 59, 59, 999999999, time.UTC)
    g, y, _ := GetWareki(d)
    if g != "大正" {
        t.Errorf("%s is not 大正", g)
    }
    if y != 15 {
        t.Errorf("%d is not valid Wareki year", y)
    }
}

func TestGetWareki03_01(t *testing.T) {
    d := time.Date(1926, 12, 25, 0, 0, 0, 0, time.UTC)
    g, y, _ := GetWareki(d)
    if g != "昭和" {
        t.Errorf("%s is not 昭和", g)
    }
    if y != 1 {
        t.Errorf("%d is not valid Wareki year", y)
    }
}

func TestGetWareki03_02(t *testing.T) {
    d := time.Date(1989, 1, 7, 23, 59, 59, 999999999, time.UTC)
    g, y, _ := GetWareki(d)
    if g != "昭和" {
        t.Errorf("%s is not 昭和", g)
    }
    if y != 64 {
        t.Errorf("%d is not valid Wareki year", y)
    }
}

func TestGetWareki04_01(t *testing.T) {
    d := time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC)
    g, y, _ := GetWareki(d)
    if g != "平成" {
        t.Errorf("%s is not 平成", g)
    }
    if y != 1 {
        t.Errorf("%d is not valid Wareki year", y)
    }
}

func TestGetWareki05_01(t *testing.T) {
    d := time.Date(1867, 12, 31, 23, 59, 59, 999999999, time.UTC)
    _, _, e := GetWareki(d)
    if e == nil {
        t.Errorf("%v is valid date", d)
    } else if fmt.Sprintf("%v", e) != "1867-12-31 23:59:59.999999999 +0000 UTC is invalid date" {
        t.Errorf("%v", e)
    }
}

func TestGetWareki06_01(t *testing.T) {
    d := time.Date(2012, 3, 28, 0, 0, 0, 0, time.UTC)
    w := NewWareki(d)
    if w.Gengo != "平成" {
        t.Errorf("%s is not 平成", w.Gengo)
    }
    if w.Year != 24 {
        t.Errorf("%d is not valid Wareki year", w.Year)
    }
    if w.String() != "平成24年3月28日" {
        t.Errorf("%s is not expected date", w.String())
    }
    if d != w.Date {
        t.Errorf("%v is not expected date", w)
    }
}

func TestGetWareki06_02(t *testing.T) {
    d := time.Date(1926, 12, 26, 0, 0, 0, 0, time.UTC)
    w := NewWareki(d)
    if w.Gengo != "昭和" {
        t.Errorf("%s is not 昭和", w.Gengo)
    }
    if w.Year != 1 {
        t.Errorf("%d is not valid Wareki year", w.Year)
    }
    if w.String() != "昭和元年12月26日" {
        t.Errorf("%s is not expected date", w.String())
    }
    if d != w.Date {
        t.Errorf("%v is not expected date", w)
    }
}
