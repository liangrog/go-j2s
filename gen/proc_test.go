package gen

import (
	"fmt"
	"github.com/liangrog/go-j2s/conf"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestMust(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	Must(fmt.Errorf("testing error"))
}

func TestCamelCase(t *testing.T) {
	s := "bl_ha-Ha_lu"
	assert.Equal(t, "BlHaHaLu", CamelCase(s))
}

func TestFindJsonFiles(t *testing.T) {
	// Prepare dirs
	tmpdir, _ := ioutil.TempDir("", "goj2s")
	childdir := filepath.Join(tmpdir, "child")
	os.MkdirAll(childdir, 0775)
	defer os.RemoveAll(tmpdir)

	// Prepare json files
	f1 := filepath.Join(tmpdir, "1.json")
	f2 := filepath.Join(tmpdir, "2.json")
	f3 := filepath.Join(childdir, "3.json")
	f4 := filepath.Join(childdir, "4.json")
	os.Create(f1)
	os.Create(f2)
	os.Create(f3)
	os.Create(f4)

	a := conf.NewArgs()
	a.InPath = tmpdir

	// Non-recursive
	r1, _ := findJsonFiles(a)
	assert.Equal(t, 2, len(r1))

	// Recursively search
	a.Recursive = true
	r2, _ := findJsonFiles(a)
	assert.Equal(t, 4, len(r2))

	// Test specify file
	a.InFile = []string{f4}
	r3, _ := findJsonFiles(a)
	assert.Equal(t, 1, len(r3))

	// Test excluded paths
	a.InFile = []string{}
	a.ExclPath = []string{"child"}
	r4, _ := findJsonFiles(a)
	assert.Equal(t, 2, len(r4))
}

func TestGetKind(t *testing.T) {
	var s interface{}
	s = "stringing"

	rt := reflect.TypeOf(s)
	assert.Equal(t, "string", getKind(rt).String())

	s = nil
	rt = reflect.TypeOf(s)
	assert.Equal(t, "interface", getKind(rt).String())
}

func TestGenGoCodes(t *testing.T) {
	a := conf.NewArgs()
	a.Recursive = true
	_, err := genGoCodes(a)
	assert.Nil(t, err)
}

func TestWriteGoCode(t *testing.T) {
	a := conf.NewArgs()
	a.Recursive = true
	a.PkgName = "goj2s"

	c, err := genGoCodes(a)
	assert.Nil(t, err)

	tmpdir, _ := ioutil.TempDir("", "goj2s")
	defer os.RemoveAll(tmpdir)

	a.OutPath = tmpdir

	ti := time.Now()
	err = writeGoCode(a.GetOutFile(), fileLayout, ti.Format(time.RFC3339), a.PkgName, c)
	assert.Nil(t, err)
}
