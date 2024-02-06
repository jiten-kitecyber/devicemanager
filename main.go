package main

import (
	"devicemanager/dm"
	"fmt"
)

func main() {
	println("Muruga! Help me")
	globaldns := new(dm.GlobalDNS)
	fmt.Println(globaldns.GetDNS("all"))
	fmt.Println(globaldns.SetDNS("", "8.8.8.8", "8.8.4.4"))

	fmt.Println(globaldns.GetDNS("all"))

}
