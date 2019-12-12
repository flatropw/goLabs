package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	connection, _ := net.Dial("tcp", config.address+":"+config.port)
	defer connection.Close()
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Request:")

		text, err := reader.ReadString(delimiter)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		text = strings.TrimSuffix(text, string(delimiter))
		if text == config.exitPhrase {
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
