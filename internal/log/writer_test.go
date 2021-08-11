package log

import (
	"bytes"
	"io"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestWriter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		give []string
		want string
	}{
		{desc: "empty"},
		{
			desc: "split message",
			give: []string{"foo\nbar"},
			want: unlines(
				"[x] foo",
				"[x] bar",
			),
		},
		{
			desc: "ends with a newline",
			give: []string{"foo\n", "bar\n"},
			want: unlines(
				"[x] foo",
				"[x] bar",
			),
		},
		{
			desc: "no newlines",
			give: []string{"foo", "bar"},
			want: unlines(
				"[x] foobar",
			),
		},
		{
			desc: "newline late",
			give: []string{"foo", "b\nar"},
			want: unlines(
				"[x] foob",
				"[x] ar",
			),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			var buff bytes.Buffer
			log := New(&buff).WithName("x")

			w := Writer{Log: log}
			for _, s := range tt.give {
				io.WriteString(&w, s)
			}
			td.CmpNoError(t, w.Close())

			td.Cmp(t, buff.String(), tt.want)
		})
	}
}
