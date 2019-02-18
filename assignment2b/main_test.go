package main

import "testing"

func TestMain(t *testing.T) {
	filename := "_testingFile"
	bs := make([]byte, 99999)
	readFile(filename, bs)
	if bs[0] != 'B' {
		t.Error("First byte of the byte slice should be 102 (letter 'B') but it was ", bs[0])
	}
}