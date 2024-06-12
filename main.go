package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	// the expression len(os.Args) evaluates to the number of arguments passed to the program when it was executed.
	// This includes the program name itself as the first argument and any additional arguments provided.
	// len: This is a built-in function that returns the length of a collection, such as a slice or array.
	//os.Args: This is a variable that holds a slice of strings representing the arguments passed to the program.

	if len(os.Args) < 2 {

		printHelp()

		os.Exit(0)
	}

	function := os.Args[1]

	switch function {

	case "help":

		printHelp()

	case "Encrypt":

		encryptHandle()

	case "decrypt":

		decryptHandle()

	default:

		fmt.Println("Run encryt to encrypt a file ,And Run decrypt to decrypta file")

		os.Exit(1)

		// here we  use os.Exit function because it is  used to terminate the currently running program immediately.

	}

}

func printHelp() {

	fmt.Println("File encryption ")

	fmt.Println("Simple file encryption for day to day life need")

	fmt.Println("")

	fmt.Println("Usage :")

	fmt.Println("")

	fmt.Println("go run . encrypt D:/projects/go_encryption ")

	fmt.Println("")

	fmt.Println("Command :")

	fmt.Println("")

	fmt.Println("\t encrypt \t encrypt a file to given a password")

	fmt.Println("\t decrypt \t decrypt a file using password")

	fmt.Println("\t help \t\t Display Help ")

}

// here we use this function to encryption of a file

func encryptHandle() {

	if len(os.Args) < 3 {

		fmt.Println("missing the path to the file. For more info run . help")

		os.Exit(0)

	}
	// := is a shorthand for variable declaration and assignment within functions.
	//It automatically infers the type from the assigned value, making it concise and readable.

	file := os.Args[2]

	if !validateFile(file) {

		panic("File not found")

	}

	password := getPassword()

	fmt.Println("\nEncrypting...")

	filecrypt.Encrypt(file, password)

	fmt.Println("\nFile is sucessfully encrypted.")
}

// here we use this function to decryption of a file
func decryptHandle() {

	if len(os.Args) < 3 {

		fmt.Println("missing the path to the file. For more info run go run . help")

		os.Exit(0)

	}

	file := os.Args[2]

	if !validateFile(file) {

		panic("File not found")

	}

	fmt.Print("Enter Password")

	password, _ := term.ReadPassword(0)

	fmt.Println("\nDecrypting...")

	filecrypt.Decrypt(file, password)

	fmt.Println("\nFile is sucessfully decrypted.")
}

// here we use this function to get a password of a file  a file

func getPassword() []byte {

	fmt.Print("Enter Password :- ")

	password, _ := term.ReadPassword(0)

	// the term package provides a way to interact with the terminal.
	//If you want to read a password from the terminal without echoing it back (for security reasons),
	// you can use the term.ReadPassword function.

	fmt.Print("\nConfirm Password")

	password2, _ := term.ReadPassword(0)

	if !validatePassword(password, password2) {

		fmt.Println("\nPasswords do not match ,Please try again")

		return getPassword()
	}

	return password

}

// here we use this function to validate a password

func validatePassword(password []byte, password2 []byte) bool {

	if !bytes.Equal(password, password2) {

		return false
	}

	return true

}

// here we use this to validate a file

func validateFile(file string) bool {

	if _, err := os.Stat(file); os.IsNotExist(err) {

		return false

	}

	return true

}
