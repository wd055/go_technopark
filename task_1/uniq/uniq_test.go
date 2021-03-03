package uniq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniq(t *testing.T) {
	testCases := []struct {
		name string
		in   []string
		opt  FlagsStruct
		out  []string
	}{
		{
			name: "Без параметров",
			in:   []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			opt:  FlagsStruct{},
			out:  []string{"I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
		},
		{
			name: "Параметры: С",
			in:   []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			opt:  FlagsStruct{C: true},
			out:  []string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."},
		},
		{
			name: "Параметры: D",
			in:   []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			opt:  FlagsStruct{D: true},
			out:  []string{"I love music.", "I love music of Kartik.", "I love music of Kartik."},
		},
		{
			name: "Параметры: U",
			in:   []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			opt:  FlagsStruct{U: true},
			out:  []string{"", "Thanks."},
		},
		{
			name: "Параметры: I",
			in:   []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of Kartik.", "I love MuSIC of Kartik."},
			opt:  FlagsStruct{I: true},
			out:  []string{"I LOVE MUSIC.", "I love MuSIC of Kartik.", "Thanks.", "I love music of Kartik."},
		},
		{
			name: "Параметры: F",
			in:   []string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{F: 1},
			out:  []string{"We love music.", "", "I love music of Kartik.", "Thanks."},
		},
		{
			name: "Параметры: S",
			in:   []string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{S: 1},
			out:  []string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.out, Uniq(testCase.in, testCase.opt))
		})
	}
}
