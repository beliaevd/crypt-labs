package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func TextEncrypt(str string, Key int) string {
	var cipherText bytes.Buffer
	for _, char := range str {
		if char == 32 {
			cipherText.WriteByte(byte(char))
			continue
		}

		var charPos = GetCharPos(char)
		var cipherChar = GetChar((charPos+Key)%26, char)
		cipherText.WriteString(cipherChar)
	}

	return cipherText.String()
}

func TextDecrypt(str string, Key int) string {
	var cipherText bytes.Buffer

	for _, char := range str {

		if char == 32 {
			cipherText.WriteByte(byte(char))
			continue
		}

		var charPos = GetCharPos(char)
		var mod26 = (charPos - Key) % 26
		if mod26 < 0 {
			mod26 += 26
		}
		var origChar = GetChar(mod26, char)
		cipherText.WriteString(origChar)
	}

	return cipherText.String()
}

func GetChar(alphabetPos int, asciiVal int32) string {
	var isLowercase = asciiVal >= 97
	if isLowercase {
		return string('a' - 1 + alphabetPos)
	}
	return string('A' - 1 + alphabetPos)
}

func GetCharPos(asciiVal int32) int {

	asciiVal |= 32
	return int(asciiVal) - 96
}

func main() {
	text, err := ioutil.ReadFile("analysis/text.txt")
	if err != nil {
		panic(err)
	}
	string := string(text)
	
	TextEncrypt := TextEncrypt(string, 3)
	TextDecrypt := TextDecrypt(TextEncrypt, 3)
	fmt.Println(TextEncrypt)

	fmt.Println("\n ------------")
	fmt.Println(TextDecrypt)

}
