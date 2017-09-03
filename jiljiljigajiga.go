package main

import "./core"

import "fmt"

func main() {
	c := make(chan bool)
	go core.ListenToServer(c)
	<-c
	fmt.Println("Stopping server")
}
