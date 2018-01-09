package store

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMmapedfile(t *testing.T) {
	fn := "test"
	m, err := newMappedfile(fn, os.Getpagesize())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.Remove(fn)
		m.close()
	}()

	m.Write([]byte("test"))
	err = m.flush()
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile(fn)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", string(b))
}
