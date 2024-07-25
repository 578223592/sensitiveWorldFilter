package wordsFilter

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
	wf := NewWordsFilter()
	wf.RemoveSensitiveWords("shif")
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

func TestWordsFilter_FilterAll(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{text: "坏蛋-3-441大坏蛋"},
			want: "-3-441",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wf := NewWordsFilter()
			wf.Add("坏蛋", "大坏蛋")
			if got := wf.FilterAll(tt.args.text); got != tt.want {
				t.Errorf("FilterAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
