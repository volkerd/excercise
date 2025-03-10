package main

import (
	"fmt"
	"github.com/volkerd/exercise/cmd"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	cmd.Exec()
}
