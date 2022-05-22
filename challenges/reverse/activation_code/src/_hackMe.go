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
	c = []byte{ /* This is a secret */ }
	key = []byte{ /* This is a secret */ }
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your activation code: ")
	activationCode, _ := reader.ReadString('\n')
	activationCode = strings.Replace(activationCode, "\n", "", -1)
	// for windows \r\n
	activationCode = strings.Replace(activationCode, "\r", "", -1)

	if activationCode == "This is a fake activation code!" {
		fmt.Println("Congratulations!")
		fmt.Println("This is your flag:", getFlag())
	} else {
		fmt.Println("Incorrect activation code! QQ")
	}
}
