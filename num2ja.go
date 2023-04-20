package num2ja

import (
	"bytes"
	"fmt"
)

var (
	knum = []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	ksub = []string{"", "十", "百", "千"}
	kpar = []string{
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

var (
	hnum = []map[rune]string{
		{'0': "ぜろ"},
		{'0': "いち"},
		{'0': "に"},
		{'0': "さん"},
		{'0': "よん"},
		{'0': "ご"},
		{'0': "ろく", '3': "ろっ"},
		{'0': "なな"},
		{'0': "はち", '3': "はっ", '4': "はっ"},
		{'0': "きゅう"},
	}
	hdig = []map[rune]string{
		{'0': ""},
		{'0': "まん"},
		{'0': "おく"},
		{'0': "ちょう"},
		{'0': "けい"},
		{'0': "がい"},
		{'0': "じょ"},
		{'0': "じょう"},
		{'0': "こう"},
		{'0': "かん"},
		{'0': "せい"},
		{'0': "さい"},
		{'0': "ごく"},
		{'0': "こうがしゃ"},
		{'0': "あそうぎ"},
		{'0': "なゆた"},
		{'0': "ふかしぎ"},
		{'0': "むりょうたいすう"},
	}
	hndg = []map[rune]string{
		{'0': ""},
		{'0': "じゅう"},
		{'0': "ひゃく", '3': "ぴゃく", '6': "ぴゃく", '8': "ぴゃく"},
		{'0': "せん", 3: "ぜん"},
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
		if rs[j] != '0' {
			return true
		}
	}
	return false
}

func ToKanji(i int64) string {
	s := fmt.Sprint(i)
	var buf bytes.Buffer
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	for i, r := range rs {
		c := r - 48
		if i%4 == 0 && i > 0 && partsIn(rs, i) {
			buf.WriteString(kpar[i/4])
		}
		if c != 0 {
			buf.WriteString(ksub[i%4])
		}
		if !(i%4 != 0 && c == 1) {
			buf.WriteString(knum[c])
		}
	}
	rs = []rune(buf.String())
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func hasKey(m map[rune]string, k rune) bool {
	_, ok := m[k]
	return ok
}

func ToHira(i int64) string {
	if i == 0 {
		return hnum[0][0]
	}
	rs := []rune(fmt.Sprint(i))
	ln := len(rs)
	if ln >= 17*4+1 {
		return hdig[17][0]
	}
	ret := ""

	for n := 0; n < ln; n++ {
		if rs[n] != '0' && (rs[n] != '1' || (ln-n)%4 == 1) {
			if hasKey(hnum[rs[n]-'0'], rune(ln-n+'0')) {
				ret += hnum[rs[n]-'0'][rune(ln-n+'0')]
			} else {
				ret += hnum[rs[n]-'0']['0']
			}
		}
		if rs[n] != '0' {
			if hasKey(hndg[(ln-n-1)%4], rs[n]) {
				ret += hndg[(ln-n-1)%4][rs[n]]
			} else {
				ret += hndg[(ln-n-1)%4]['0']
			}

			if hasKey(hdig[(ln-n-1)/4], rs[n]) {
				ret += hdig[(ln-n-1)/4][rs[n]]
			} else {
				ret += hdig[(ln-n-1)/4]['0']
			}
		}
	}

	return ret
}
