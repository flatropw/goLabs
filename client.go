package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	connection, err := net.Dial("tcp", config.address+":"+config.port)
	if err != nil {
		fmt.Println(err)
	}
	defer connection.Close()

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Request:")

		text, err := reader.ReadString(delimiter)
		if err != nil {
				fmt.Println(err)
			break
		}

		text = strings.TrimSuffix(text, string(delimiter))
		if text == exitPhrase {
			break
		}

		_, err = fmt.Fprintf(connection, text+string(delimiter))

		if err != nil {
			fmt.Println(err)
		}

		message, err := bufio.NewReader(connection).ReadString(delimiter)

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Print("Response: " + message)
	}
}
