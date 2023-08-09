package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/AxterDoesCode/monkey/repl"
)

func main() {
    user, err := user.Current()
     
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello this is the Monkey programming language, %s", user)
    fmt.Printf("Feel free to type in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
