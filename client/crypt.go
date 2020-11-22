package main

import (
	"encoding/base64"
	"crypto/aes"
	"encoding/hex"
)

func Base64Encode(input string) string{ 
	return string(base64.StdEncoding.EncodeToString([]byte(input)))
}

func Base64Decode(input string) string{
	ou, err := base64.StdEncoding.DecodeString(input)
	if err != nil { 
		panic(err)
	}
	return string(ou)
}


func AESEncrypt(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	if err != nil { 
		panic(err)
	}
	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}


func AESDecrypt(key []byte, ct string) string{
    ciphertext, _ := hex.DecodeString(ct)
    c, err := aes.NewCipher(key)
    if err != nil { 
		panic(err)
	}
 
    pt := make([]byte, len(ciphertext))
    c.Decrypt(pt, ciphertext)
 
    s := string(pt[:])
    return s
}

func CrossHatchEncrypt(key []byte, plaintext string) string { 
	encode := Base64Encode(string(key))
	return AESEncrypt(key, encode)
}

func CrossHatchDecrypt(key []byte, ciphertext string) string { 
	decrypt := AESDecrypt(key, ciphertext)
	return Base64Decode(decrypt)
}