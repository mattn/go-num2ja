package int2kanji

import (
	"bytes"
	"fmt"
)

var (
	nums  = []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	subs  = []string{"", "十", "百", "千"}
	parts = []string{
		"",
		"万",
		"億",
		"兆",
		"京",
		"垓",
		"𥝱",
		"穣",
		"溝",
		"澗",
		"正",
		"載",
		"極",
		"恒河沙",
		"阿僧祇",
		"那由他",
		"不可思議",
		"無量大数",
	}
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func partsIn(rs []rune, i int) bool {
	for j := i; j < min(i+4, len(rs)); j++ {
		if rs[i] != '0' {
			return true
		}
	}
	return false
}

func Int2Kanji(i int64) string {
	s := fmt.Sprint(i)
	var buf bytes.Buffer
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	for i, r := range rs {
		c := r - 48
		if i%4 == 0 && i > 0 && partsIn(rs, i) {
			buf.WriteString(parts[i/4])
		}
		if c != 0 {
			buf.WriteString(subs[i%4])
		}
		if !(i%4 != 0 && c == 1) {
			buf.WriteString(nums[c])
		}
	}
	rs = []rune(buf.String())
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
