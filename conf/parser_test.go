package conf

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	os.Args = []string{
		"noop",
		"-in=tmp",
		"-excl=tmp/excl1",
		"-r=true",
		"-excl=tmp/excl2",
	}
	a := Parse()

	assert.Equal(t, "tmp", a.InPath, "inPath is not correct")
	assert.Equal(t, 2, len(a.ExclPath), "exclPath is not having the right number of argument")
	assert.Equal(t, true, a.Recursive, "recursive is false rather than true")
}
