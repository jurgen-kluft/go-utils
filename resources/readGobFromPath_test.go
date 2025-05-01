package resources_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jurgen-kluft/go-utils/assert"
	"github.com/jurgen-kluft/go-utils/resources"
)

func TestReadGobFromPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.gob")
	_ = os.Remove("testdata")
	type Data struct{ Name string }
	_ = resources.WriteGobToPath(Data{Name: "foo"}, path)
	got, err := resources.ReadGobFromPath[Data](path)
	assert.That("err should be nil", t, err, nil)
	assert.That("Name should be foo", t, got.Name, "foo")
}
