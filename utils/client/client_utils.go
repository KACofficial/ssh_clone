package client_utils

import (
	"os"
	"io"
	"fmt"
	"net"
	"bufio"
	"strings"
	"github.com/chzyer/readline"
)

func GetKey(filename string) (string, error) {
	key, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading key:", err)
        return "", err
	}
	return string(key), nil
}

func formatAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func HandleConnection(conn net.Conn, keyfile string) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    
    key, err := GetKey(keyfile)
    if err != nil {
        fmt.Println("Error getting key:", err)
        return
    }
    conn.Write([]byte(key + "\n")) // Send key to server

    k_resp, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading message:", err)
        return
    }
    k_resp = strings.TrimSpace(k_resp)
    switch k_resp {
    case "OK":
        fmt.Println("Key verified")
    case "FAIL":
        fmt.Println("Invalid key")
        return
    }
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "GSH> ",
	})
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		return
	}
	defer rl.Close()

    for {
        cmd_line, err := rl.Readline()
		if err != nil { // io.EOF or readline.ErrInterrupt
			if err == readline.ErrInterrupt {
				if len(cmd_line) == 0 {
					break
				} else {
					continue
				}
			} else if err == io.EOF {
				break
			}
			fmt.Println("Error reading command:", err)
			return
		}
		cmd_line = strings.TrimSpace(cmd_line)
		if cmd_line == "" {
			continue
		}
		conn.Write([]byte(cmd_line + "\n"))

		resp, err := reader.ReadString('\x00')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Server closed the connection.")
				return
			}
			fmt.Println("Error reading response:", err)
			return
		}
		fmt.Print(resp)
	}
}

func StartClient(address string, port int, keyfile string) {
	conn, err := net.Dial("tcp", formatAddress(address, port))
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	HandleConnection(conn, keyfile)
}