package wareki

import (
	"errors"
	"fmt"
	"time"
)

type Wareki struct {
	Gengo string
	Year  int
	Date  time.Time
}

func GetWareki(d time.Time) (string, int, error) {
	meijiBeginsOn := time.Date(1868, 1, 25, 0, 0, 0, 0, time.UTC)
	taishoBeginsOn := time.Date(1912, 7, 30, 0, 0, 0, 0, time.UTC)
	showaBeginsOn := time.Date(1926, 12, 25, 0, 0, 0, 0, time.UTC)
	heiseiBeginsOn := time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC)
	if d.Before(meijiBeginsOn) {
		return "", -1, errors.New(fmt.Sprintf("%v is invalid date", d))
	}
	var gengo string
	var base int
	switch {
	case (d.Equal(meijiBeginsOn) || d.After(meijiBeginsOn)) && d.Before(taishoBeginsOn):
		gengo = "明治"
		base = 1868
	case (d.Equal(taishoBeginsOn) || d.After(taishoBeginsOn)) && d.Before(showaBeginsOn):
		gengo = "大正"
		base = 1912
	case (d.Equal(showaBeginsOn) || d.After(showaBeginsOn)) && d.Before(heiseiBeginsOn):
		gengo = "昭和"
		base = 1926
	case d.Equal(heiseiBeginsOn) || d.After(heiseiBeginsOn):
		gengo = "平成"
		base = 1989
	}
	year := d.Year() - base + 1
	return gengo, year, nil
}

func NewWareki(d time.Time) *Wareki {
	g, y, err := GetWareki(d)
	if err != nil {
		return &Wareki{}
	}
	return &Wareki{Gengo: g, Year: y, Date: d}
}

func (w *Wareki) String() string {
	if w.Year == 1 {
		return fmt.Sprintf("%s%s年%d月%d日", w.Gengo, "元", w.Date.Month(), w.Date.Day())
	} else {
		return fmt.Sprintf("%s%d年%d月%d日", w.Gengo, w.Year, w.Date.Month(), w.Date.Day())
	}
}
