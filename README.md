# Go SSH clone
GSH(Go Shell) is purposefully missing the secure part because,  
this hasnt been tested for security
## How to build
- Clone this repo
- Run `git init <NAME>` replace `<NAME>` with the name of the package, eg `gsh`.
- Run `go build` or `go run main.go`.
## Usage
- Run it with `connect` to connect to the server.
- Run it with `host` to host the server.
- Use `-a`/`-address` to specify the address of the server(default is `localhost`).
- Use `-p`/`-port` to specify the port of the server(default is `1234`).
- Use `-k`/`-key` to specify the keyfile(default is `key.txt`).
- Host and client MUST have the same value in the `key.txt` file.
