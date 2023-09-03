package main

import (
	"fmt"
	"net"
)

func main() {

	_, err := net.Dial("tcp", "whitebox.sysadmins.help:2")

	if err == nil {
		fmt.Println("connection sucessfull")
	} else {
		fmt.Println(err)
	}

}
