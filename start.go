package main

import (
	"fmt"
	"os"

	"github.com/wedojava/go-sshmitm/services"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "go" {
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
	if len(os.Args) > 1 && os.Args[1] == "please" {
		fmt.Println("How's the fucking arguments going?")
		fmt.Println("go: run at port 22")
		fmt.Println("go 33: run at port 33")
		fmt.Println("go 22 ./abc: run at port 22 with ./abc as private key.")
	}
	if len(os.Args) == 1 {
		fmt.Print("Fuck u jiu zai tomorrow!")
	}
	//services.LocalForward()
}
