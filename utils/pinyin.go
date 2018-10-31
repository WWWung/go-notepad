package utils

import (
	"github.com/mozillazg/go-pinyin"
)

//ToPinYin1 ..
func ToPinYin1(h string) (py1 string, py2 string) {
	a := pinyin.NewArgs()
	s := pinyin.LazyPinyin(h, a)
	for _, item := range s {
		py1 += item
		py2 += Substr(item, 0, 1)
	}
	return
}

//Substr 截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
