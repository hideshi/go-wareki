package gengo

import (
    "testing"
    "time"
    "fmt"
)

func TestGetGengo01_01(t *testing.T) {
    d := time.Date(1868, 1, 25, 0, 0, 0, 0, time.UTC)
    g, gy, _ := GetGengo(d)
    if g != "明治" {
        t.Errorf("%s is not 明治", g)
    }
    if gy != 1 {
        t.Errorf("%d is not valid Gengo year", gy)
    }
}

func TestGetGengo01_02(t *testing.T) {
    d := time.Date(1912, 7, 29, 23, 59, 59, 999999999, time.UTC)
    g, gy, _ := GetGengo(d)
    if g != "明治" {
        t.Errorf("%s is not 明治", g)
    }
    if gy != 45 {
        t.Errorf("%d is not valid Gengo year", gy)
    }
}

func TestGetGengo02_01(t *testing.T) {
    d := time.Date(1912, 7, 30, 0, 0, 0, 0, time.UTC)
    g, gy, _ := GetGengo(d)
    if g != "大正" {
        t.Errorf("%s is not 大正", g)
    }
    if gy != 1 {
        t.Errorf("%d is not valid Gengo year", gy)
    }
}

func TestGetGengo02_02(t *testing.T) {
    d := time.Date(1926, 12, 24, 23, 59, 59, 999999999, time.UTC)
    g, gy, _ := GetGengo(d)
    if g != "大正" {
        t.Errorf("%s is not 大正", g)
    }
    if gy != 15 {
        t.Errorf("%d is not valid Gengo year", gy)
    }
}

func TestGetGengo03_01(t *testing.T) {
    d := time.Date(1926, 12, 25, 0, 0, 0, 0, time.UTC)
    g, gy, _ := GetGengo(d)
    if g != "昭和" {
        t.Errorf("%s is not 昭和", g)
    }
    if gy != 1 {
        t.Errorf("%d is not valid Gengo year", gy)
    }
}

func TestGetGengo03_02(t *testing.T) {
    d := time.Date(1989, 1, 7, 23, 59, 59, 999999999, time.UTC)
    g, gy, _ := GetGengo(d)
    if g != "昭和" {
        t.Errorf("%s is not 昭和", g)
    }
    if gy != 64 {
        t.Errorf("%d is not valid Gengo year", gy)
    }
}

func TestGetGengo04_01(t *testing.T) {
    d := time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC)
    g, gy, _ := GetGengo(d)
    if g != "平成" {
        t.Errorf("%s is not 平成", g)
    }
    if gy != 1 {
        t.Errorf("%d is not valid Gengo year", gy)
    }
}

func TestGetGengo05_01(t *testing.T) {
    d := time.Date(1867, 12, 31, 23, 59, 59, 999999999, time.UTC)
    _, _, e := GetGengo(d)
    if e == nil {
        t.Errorf("%v is valid date", d)
    } else if fmt.Sprintf("%v", e) != "1867-12-31 23:59:59.999999999 +0000 UTC is invalid date" {
        t.Errorf("%v", e)
    }
}

func TestGetGengo06_01(t *testing.T) {
    d := time.Date(1980, 1, 30, 10, 20, 30, 0, time.UTC)
    g := NewGengo(d)
    if g.gengo != "昭和" {
        t.Errorf("%s is not 昭和", g.gengo)
    }
    if g.year != 55 {
        t.Errorf("%d is not valid Gengo year", g.year)
    }
    if g.String() != "昭和55年1月30日" {
        t.Errorf("%s is not expected date", g.String())
    }
    if d != g.date {
        t.Errorf("%v is not expected date", g)
    }
}
