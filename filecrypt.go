package filecrypt

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)
	

func Encrypt(source string, _ []byte) {

	if err, _ := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}
// here we open the source file
	srcFile, err := os.Open(source)
	if err != nil {
		panic(err.Error())

	}
//here we close the source file

defer srcFile.Close()

	// Read the plain text from source file
	
	plaintext, err := io.ReadAll(srcFile)
	if err!= nil {
        panic(err.Error())
    }

    key := password 
	 
	// here we randomize the nonce

	nonce := make([]byte,12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
    }
// password based key derivation system :- a function in which key is genrated from password  can be used as an  encryption key or as a hash value 	

// 4096 we use the iteration counter

dk := pbkdf2.Key(key , nonce,4096,32,sha1.New )
// dk - derived key

// aes - advance Encryption Standard , solid encryption standard
block, err := aes.NewCipher(dk)
 if err != nil {
	panic(err.Error())

}






func Decrypt() {

}