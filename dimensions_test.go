package gridt

import "testing"

var d, _ = NewWithCells(TopToBottom, "  ", "sdfghjuytrfdcvg", "fghuytrfdjhjkloiuygfd", "fvbhjkilkjhgfdsxcvghjui", "ooiuygtfdcvbhj", "lokjh", "gfdcfghjkiloç.,mnjhgfrdesxc", "vghjkll", "kjhgfdcvbhjkl", "kijuhygtrfdcvg", "hjklkjuyhgtfdcfvg", "hjkilkiuyt", "rd", "iuytred", "fghjklkj", "hgfrdesxdesxcvfgtfcvbhjui", "jhkloiujuytgftrewsdewsxcfgtgvbnji", "jkoiujhytredewazxdrfdcvbhuhnmklo").FitIntoColumns(uint(len("jkoiujhytredewazxdrfdcvbhuhnmklo") * 2))

func BenchmarkOldString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d.OldString()
	}
}
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d.String()
	}
}
