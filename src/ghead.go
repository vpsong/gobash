package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	numbPtr := flag.Int("n", 10, "int")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("no file")
		return
	}
	head(flag.Args()[0], *numbPtr)
}

func head(file string, numbPtr int) {
	fmt.Println(file)
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Print("error")
		return
	}
	buff := bufio.NewReader(f)
	lineNum := 0
	for {
		if numbPtr <= lineNum {
			break
		}
		lineNum++
		line, err := buff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}
}
