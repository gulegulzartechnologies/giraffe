package giraffe

import (
	"crypto/md5"
	"fmt"
)

func Hasher(hashstring string) string {
	hash := md5.Sum([]byte(hashstring))
	return fmt.Sprintf("%x", hash)
}
