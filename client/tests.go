//only in the package for testing methods, pyloads, etc.

package main

import (
	"fmt"
)

func main() { 
	//test encryption
	var enc string
	fmt.Println("testing base 64 encode: ", Base64Encode("thequickfoxjumpedoverthelazydog"))
	enc = Base64Encode("thequickfoxjumpedoverthelazydog")
	fmt.Println("testing base 64 decode: ", Base64Decode(enc))
	fmt.Println("testing AES encrypt: ")
	// key must be 32 bit
	key := "hellothisisakeyru3i4lfk4i39d0323"
	//plaintext needs to be at least 16 bits
	secret := "this ismy secrett"

	crypton := AESEncrypt([]byte(key),  secret)
	fmt.Println(crypton)
	fmt.Println("testing AES decrypt: ", AESDecrypt([]byte(key), crypton))



}