package main

import (
	"flag"
	"fmt"
	"os"
	cl "ssh/utils/client"
	svr "ssh/utils/server"
)

func main() {
	var (
		port    int
		address string
		keyfile string
	)

	flag.StringVar(&address, "address", "127.0.0.1", "server address")
	flag.StringVar(&address, "a", "127.0.0.1", "server address (shorthand)")
	flag.IntVar(&port, "port", 1234, "server port")
	flag.IntVar(&port, "p", 1234, "server port (shorthand)")
	flag.StringVar(&keyfile, "key", "key.txt", "keyfile")
	flag.StringVar(&keyfile, "k", "key.txt", "keyfile (shorthand)")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Please specify 'connect' or 'host' command")
		os.Exit(1)
	}
	command := flag.Arg(0)

	fmt.Println("Running with these args:")
	fmt.Println("  port:", port)
	fmt.Println("  address:", address)
	fmt.Println("  keyfile:", keyfile)
	fmt.Println("  type:", command)

	
	switch command {
	case "connect":
		cl.StartClient(address, port, keyfile)
	case "host":
		svr.StartServer(address, port, keyfile)
	default:
		fmt.Println("Invalid command:", command)
		os.Exit(1)
	}
}
