package sentive_word_filter

import (
	"testing"
)

func TestWordsFilter(t *testing.T) {
	texts := []string{
		"坏蛋",
		"三个坏蛋",
		"妲己",
		"アンジェラ",
		"ความรุ่งโรจน์",
	}
	wf := NewWordsFilter(DefaultIgnoreRunes, true)
	wf.RemoveSensitiveWord("shif")
	wf.Add(texts...)

	ok, word := wf.IsContainsSensitiveWord("都是fdsafa坏蛋，，，")
	if !ok {
		t.Error("Test_IsContainsSensitiveWord error")
	} else {
		t.Log(word)
	}

}

func Test_stripSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test_stripSpace", args: args{str: "   "}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripSpace(tt.args.str); got != tt.want {
				t.Errorf("stripSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
