package pkg

import (
	"fmt"
	"math/rand"
)

func GenerateUUID(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
