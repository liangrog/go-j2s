package conf

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	a := NewArgs()
	assert.Equal(t, "ptr", reflect.TypeOf(a).Kind().String(), "New argument is not a pointer")
	assert.Equal(t, "string", reflect.TypeOf(a.InPath).Kind().String(), "inPath is not string")
	assert.Equal(t, "bool", reflect.TypeOf(a.Recursive).Kind().String(), "recursive is not bool")
	assert.Equal(t, 0, len(a.ExclPath), "exclPath is not zero length")
	assert.Equal(t, true, reflect.DeepEqual(a.ExclPath, MultiStrVar{}), "exclPath is not type MultiStrVar")
}

func TestValidate(t *testing.T) {
	a := NewArgs()
	assert.Nil(t, a.Validate())

	a.PkgName = ""
	assert.NotNil(t, a.Validate())
}
