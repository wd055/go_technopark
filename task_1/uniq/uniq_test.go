package uniq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	testCases := []struct {
		name string
		in   []string
		opt  FlagsStruct
		out  []string
	}{
		{
			name: "test1",
			in:   []string{"I love music.", "I love music.", "I love music.", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{},
			out:  []string{"I love music.", "I love music of Kartik.", "Thanks."},
		},
		{
			name: "test2",
			in:   []string{"I love music.", "I love music.", "I love music.", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{C: true},
			out:  []string{"3 I love music.", "2 I love music of Kartik.", "1 Thanks."},
		},
		{
			name: "test3",
			in:   []string{"I love music.", "I love music.", "I love music.", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{D: true},
			out:  []string{"I love music.", "I love music of Kartik."},
		},
		{
			name: "test4",
			in:   []string{"I love music.", "I love music.", "I love music.", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{U: true},
			out:  []string{"Thanks."},
		},
		{
			name: "test5",
			in:   []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks."},
			opt:  FlagsStruct{I: true},
			out:  []string{"I LOVE MUSIC.", "I love MuSIC of Kartik.", "Thanks."},
		},
		{
			name: "test6",
			in:   []string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{F: 1},
			out:  []string{"We love music.", "I love music of Kartik.", "Thanks."},
		},
		{
			name: "test7",
			in:   []string{"I love music.", "A love music.", "C love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{S: 1},
			out:  []string{"I love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		},
		{
			name: "test8",
			in:   []string{"I love music.", "A love music.", "C love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			opt:  FlagsStruct{C: true, D: true, U: true},
			out:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.out, Uniq(testCase.in, testCase.opt))
		})
	}
}
