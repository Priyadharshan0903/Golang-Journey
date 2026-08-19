package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {

	payload := []byte("hello high value software engineer")
	hashAndBroadcast(bytes.NewReader(payload))

}

func hashAndBroadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	// here the reader is already read which means emptied
	// water is already drunken
	// no more thing left in the reader
	// so that's why we need to go with composability
	if err != nil {
		return err
	}

	hash := sha1.Sum(b)
	fmt.Println(hex.EncodeToString(hash[:]))

	return broadCast(r)
}

func broadCast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string of the bytes: ", string(b))

	return nil
}
