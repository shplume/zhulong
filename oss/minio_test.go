package oss

import (
	"os"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	err := CreateBucket("")
	if err != nil {
		t.Error(err)
	}
}

func TestUploadFile(t *testing.T) {
	name := "init.go"
	file, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}

	info, _ := file.Stat()

	i, err := UploadFile("zhulong", name, file, info.Size())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(i)
}

func TestGetFileTemporaryURL(t *testing.T) {
	url, err := GetFileTemporaryURL("zhulong", "init.go")
	if err != nil {
		t.Fatal(url)
	}

	t.Log(url)
}
