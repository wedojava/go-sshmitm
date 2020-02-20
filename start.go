package main

import (
	"fmt"
	"os"

	"github.com/wedojava/go-sshmitm/services"
)

const KEYWORD = "kolin"

func main() {
	fuckOff()
	switch os.Args[1] {
	case "please":
		usage()
	case KEYWORD:
		process()
	}
}

func fuckOff() {
	if len(os.Args) == 1 {
		fmt.Print("Fuck u jiu zai tomorrow!")
		os.Exit(1)
	}
}

func process() {
	port := "22"
	rsa := "./id_rsa"
	for i, s := range os.Args[1:] {
		if i == 1 {
			port = s
		}
		if i == 2 {
			rsa = s
		}
	}
	services.Server(rsa, port)
}

func usage() {
	fmt.Println("How's the fucking arguments going?")
	fmt.Println(KEYWORD + ": run at port 22")
	fmt.Println(KEYWORD + " 33: run at port 33")
	fmt.Println(KEYWORD + " 22 ./abc: run at port 22 with ./abc as private key.")
}
