package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"time"
)

func uploadImg() {
	h := sha1.New()
	h.Write([]byte("ExampleHttpString"))
	sha1HttpString := h.Sum(nil)

	var hashFunc = sha1.New
	h = hmac.New(hashFunc, []byte("KXasxpTyHMtADUXBZsseS1aqusZ1Wy5w"))
	t := time.Now()

	h.Write([]byte(fmt.Print(t.Unix())))
	signKey := h.Sum(nil)
}