package main

import (
	"os"
	"log"
)

func getHostname() string{
    name, err := os.Hostname()
    if err != nil { 
        panic(err)
    }
    return name
}

func getOS() string{ 
	var os string
	if runtime.GOOS == "windows" {
		os = "Windows"
	}else if runtime.GOOS == "linux" { 
		os = "Linux"
	}
	return os

}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}