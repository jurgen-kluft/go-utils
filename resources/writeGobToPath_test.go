package resources_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jurgen-kluft/go-utils/assert"
	"github.com/jurgen-kluft/go-utils/resources"
)

func TestWriteGobToPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.gob")
	_ = os.Remove("testdata")
	gob := struct {
		Bar string
	}{
		Bar: "bar",
	}
	err := resources.WriteGobToPath(gob, path)
	assert.That("err should be nil", t, err, nil)
}
