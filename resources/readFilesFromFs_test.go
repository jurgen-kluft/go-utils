package resources_test

import (
	"embed"
	"testing"

	"github.com/jurgen-kluft/go-utils/assert"
	"github.com/jurgen-kluft/go-utils/resources"
)

//go:embed testembed
var testembed embed.FS

func TestReadFilesFromFs(t *testing.T) {
	filesMap, err := resources.ReadFilesFromFs(testembed)
	assert.That("err should be nil", t, err, nil)
	assert.That("filesMap should have file bar.txt", t, filesMap["testembed/bar.txt"] != nil, true)
	assert.That("file bar.txt content should be correct", t, filesMap["testembed/bar.txt"], []byte("bar"))
	assert.That("filesMap should have file foo.txt", t, filesMap["testembed/foo.txt"] != nil, true)
	assert.That("file foo.txt content should be correct", t, filesMap["testembed/foo.txt"], []byte("foo"))
}
