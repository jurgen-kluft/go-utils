package resources_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jurgen-kluft/go-utils/assert"
	"github.com/jurgen-kluft/go-utils/resources"
)

func TestReadJsonFromPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.json")
	_ = os.Remove("testdata")
	type Data struct {
		Name string `json:"name"`
	}
	_ = resources.WriteJsonToPath(Data{Name: "foo"}, path)
	got, err := resources.ReadJsonFromPath[Data](path)
	assert.That("err should be nil", t, err, nil)
	assert.That("Name should be foo", t, got.Name, "foo")
}
