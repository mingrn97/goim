package pwd

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
)

func NewPwd() {
	var hash hash.Hash = sha256.New()

	var pwd string = "admin1231q2w3e4r5t!@#$%^6ydfvdzcfgsdfdfxgo[]p[]p[]["

	len, err := hash.Write([]byte(pwd))

	if err != nil {
		fmt.Printf("pwd err: %v\n", err)
	}

	fmt.Printf("len: %d\n", len)

	var res []byte = hash.Sum(nil)

	var hexStr string = hex.EncodeToString(res)

	var b64 string = base64.StdEncoding.EncodeToString(res)

	fmt.Printf("hex: %v\n", hexStr)
	fmt.Printf("hex: %v\n", b64)
}

func Base64E() {

}
