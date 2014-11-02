package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Print("no file")
		return
	}
	cat(flag.Args()[0])
}

func cat(file string) {
	f, err := os.Open(file)
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			fmt.Print(line)
		}
	}
}
