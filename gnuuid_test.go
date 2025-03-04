package gnuuid_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/gnames/gnuuid"
)

func TestGNDomain(t *testing.T) {
	assert.Equal(t, GNDomain.String(), "90181196-fecf-5082-a4c1-411d4f314cda")
}

func TestNil(t *testing.T) {
	assert.Equal(t, Nil.String(), "00000000-0000-0000-0000-000000000000")
}

func TestNew(t *testing.T) {
	assert.Equal(t, New("Homo sapiens").String(),
		"16f235a0-e4a3-529c-9b83-bd15fe722110")
}

func TestFromFile(t *testing.T) {
	// Create a temporary file with known content
	content := []byte("test content")
	tmpfile, err := os.CreateTemp("", "example")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name()) // clean up

	_, err = tmpfile.Write(content)
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	// Generate UUID from file content
	uuid, err := FromFile(tmpfile.Name())
	assert.NoError(t, err)
	assert.Equal(t, uuid.String(), "308ee6ec-b703-5f93-9346-2c08e873f2e0")

	// Test with non-existent file
	_, err = FromFile("non_existent_file.txt")
	assert.Error(t, err)
}
