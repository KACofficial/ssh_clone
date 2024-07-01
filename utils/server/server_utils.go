package server_utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"net"
	"bufio"
)



func CheckKey(filename string, key string) (bool, error) {
    svrKey, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading key:", err)
        return false, err
    }

    return strings.TrimSpace(string(svrKey)) == strings.TrimSpace(key), nil
}

func HandleConnection(conn net.Conn, keyfile string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	key, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading key:", err)
		return
	}
	key = strings.TrimSpace(key)

	worked, err := CheckKey(keyfile, key) // verify key
	if err != nil { // handle error
		fmt.Println("Error checking key:", err)
		return
	}
	if !worked { // invalid key
		conn.Write([]byte("FAIL\n"))
        return
	}
	conn.Write([]byte("OK\n")) // key verified
	fmt.Println("Key verified")

	for {
        cmdLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command:", err)
			return
		}
		cmdLine = strings.TrimSpace(cmdLine)
		if cmdLine == "" {
			continue
		}
		if cmdLine == "exit" {
			fmt.Println("Client disconnected.")
			return
		}

		// Split command and arguments
		parts := strings.Fields(cmdLine)
		command := parts[0]
		args := parts[1:]

		cmd := exec.Command(command, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Error: %s\x00", err)))
			continue
		}

		// Send the entire output at once with a null character delimiter
		conn.Write(append(output, '\x00'))
    }
}

func formatAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func StartServer(address string, port int, keyfile string) {
	listener, err := net.Listen("tcp", formatAddress(address, port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Server listening on %s\n", formatAddress(address, port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go HandleConnection(conn, keyfile)
	}
}