package resources_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jurgen-kluft/go-utils/assert"
	"github.com/jurgen-kluft/go-utils/resources"
)

func TestWriteJsonToPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.json")
	_ = os.Remove("testdata")
	type Data struct {
		Name string `json:"name"`
	}
	err := resources.WriteJsonToPath(Data{Name: "bar"}, path)
	assert.That("err should be nil", t, err, nil)
}
