package main

import (
	"os"
	"testing"
)

func TestShred(t *testing.T) {
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
		t.Log("Shred return 1")
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
