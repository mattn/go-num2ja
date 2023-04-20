package int2kanji

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	type test struct {
		num  int64
		want string
	}
	tests := []test{
		//{num: 1, want: "一"},
		//{num: 9, want: "九"},
		//{num: 10, want: "十"},
		//{num: 11, want: "十一"},
		//{num: 21, want: "二十一"},
		//{num: 99, want: "九十九"},
		//{num: 100, want: "百"},
		//{num: 999, want: "九百九十九"},
		//{num: 1000, want: "千"},
		//{num: 9999, want: "九千九百九十九"},
		//{num: 10000, want: "一万"},
		//{num: 10020, want: "一万二十"},
		//{num: 1000020, want: "百万二十"},
		{num: 100000020, want: "一億二十"},
		//{num: 100004423, want: "一億四千四百二十三"},
		//{num: 180004423, want: "一億八千万四千四百二十三"},
	}
	for _, tt := range tests {
		got := Int2Kanji(tt.num)
		if got != tt.want {
			log.Fatalf("want %v for %v, but got %v", tt.want, tt.num, got)
		}
	}
}
