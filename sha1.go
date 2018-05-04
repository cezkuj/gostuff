package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func GetSHA1Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func test_sha1() {
	fmt.Println(GetSHA1Hash("admin"))

}
