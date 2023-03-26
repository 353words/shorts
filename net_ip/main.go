/*
#golang gem: Use the "net" package to check if IP is in a network

_, network, _ := net.ParseCIDR("193.147.0.0/16")
ip := net.ParseIP("193.147.7.247")
fmt.Println(network.Contains(ip)) // true
*/
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	_, network, err := net.ParseCIDR("193.147.0.0/16")
	if err != nil {
		log.Fatal(err)
	}
	ip := net.ParseIP("193.147.7.247")
	fmt.Println(network.Contains(ip))
}
