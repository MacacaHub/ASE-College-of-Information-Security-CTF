package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	c   []byte
	key []byte
)

func getFlag() string {
	m := make([]byte, len(c))
	for i := 0; i < len(c); i++ {
		m[i] = c[i] ^ key[i%len(key)]
	}
	flag := string(m)
	return flag
}

func init() {
	c = []byte{22, 121, 42, 39, 125, 80, 107, 61, 62, 76, 112, 76, 125, 105, 15, 15, 67, 14, 18, 55, 65, 106, 9, 85, 122, 56, 83, 39, 69, 62, 41, 106, 5, 80, 118, 120, 21, 90, 17, 115}
	key = []byte{0x50, 0x35, 0x6b, 0x60, 0x6, 0x27, 0x5b, 0x6a, 0x61, 0xe, 0x9, 0x3c, 0x49, 0x1a, 0x7c}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your activation code: ")
	activationCode, _ := reader.ReadString('\n')
	activationCode = strings.Replace(activationCode, "\n", "", -1)
	// for windows \r\n
	activationCode = strings.Replace(activationCode, "\r", "", -1)

	if activationCode == "A?cN0=G49{VPvm8ZS2_z4YyPiz99|;Y" {
		fmt.Println("Congratulations!")
		fmt.Println("This is your flag:", getFlag())
	} else {
		fmt.Println("Incorrect activation code! QQ")
	}
}
