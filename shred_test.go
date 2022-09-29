package main

import (
	"bytes"
	"os"
	"testing"
)

func TestShred(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	var testPath = "/tmp/testme"
	var file, err = os.OpenFile(testPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Log("Error creating the test file")
		t.Fail()
	}

	file.WriteString("Delete me")
	file.Sync()
	file.Close()

	var ret = shred(testPath)

	if ret != 0 {
		t.Log("Shred return 1, should be 0")
		t.Fail()
	}

	_, err = os.OpenFile(testPath, os.O_RDONLY, 0644)
	if err == nil {
		t.Fail()
	}

	if err.Error() != "open /tmp/testme: no such file or directory" {
		t.Fail()
	}
}

func TestShredNoFile(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	var testPath = "/tmp/testme"

	var ret = shred(testPath)

	if ret == 0 {
		t.Log("Shred return 0, should be 1")
		t.Fail()
	}
}
