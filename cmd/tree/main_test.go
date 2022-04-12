package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestTree(t *testing.T) {
	tests := map[string]struct {
		Path   string
		Level  int
		Expect string
	}{
		"1": {
			Path:  "./../../_example",
			Level: -1,
			Expect: `├── 01
│   ├── 02
│   │   ├── hoge.1.txt
│   │   └── hoge.txt
│   └── hoge.txt
├── 03
│   ├── 04
│   │   └── hoge.txt
│   └── hoge.txt
├── 05
│   └── hoge.txt
├── 06
│   └── hoge.txt
├── 07
│   └── hoge.txt
├── 08
│   ├── 09
│   │   ├── 10
│   │   │   ├── hoge.1.txt
│   │   │   └── hoge.txt
│   │   └── hoge.txt
│   └── 11
│       └── hoge.txt
├── 12
│   └── hoge.txt
├── 13
│   ├── 14
│   │   └── hoge.txt
│   └── hoge.txt
└── 15
    ├── 16
    └── hoge.txt
`,
		},
		"2": {
			Path:  "./../../_example",
			Level: 1,
			Expect: `├── 01
├── 03
├── 05
├── 06
├── 07
├── 08
├── 12
├── 13
└── 15
`,
		},
	}
	for _, test := range tests {
		buf := new(bytes.Buffer)
		out = buf
		t.Run("", func(t *testing.T) {
			level = test.Level
			tree(0, "", test.Path)
			str := strings.Replace(buf.String(), "\x1b[34m", "", -1)
			str = strings.Replace(str, "\x1b[0m", "", -1)
			if str != test.Expect {
				dmp := diffmatchpatch.New()
				a, b, c := dmp.DiffLinesToChars(str, test.Expect)
				diffs := dmp.DiffMain(a, b, false)
				result := dmp.DiffCharsToLines(diffs, c)
				fmt.Println(result)
				t.Errorf("")
			}
		})
	}
}
