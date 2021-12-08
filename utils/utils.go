package utils

// import (
// 	"crypto/rand"
// 	"math/big"
// 	"time"

// 	"github.com/goombaio/namegenerator"
// )

// func RandomString(n int) string {
// 	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
// 	ret := make([]byte, n)
// 	for i := 0; i < n; i++ {
// 		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
// 		if err != nil {
// 			return ""
// 		}
// 		ret[i] = letters[num.Int64()]
// 	}

// 	return string(ret)
// }

// func RandomName() string {
// 	seed := time.Now().UTC().UnixNano()
// 	nameGenerator := namegenerator.NewNameGerator(seed)
// 	name := nameGenerator.Generate()
// 	return name
// }
