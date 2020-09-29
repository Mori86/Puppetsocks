//Puppetsocks
//by Mori
package main


import (
    "fmt"
    "log"
    "net/http"
    "os"
 
    _ "strings"

)

func isRoot() bool{
        if os.Getenv("SUDO_UID") == "1000" {
            return true
        }
        return false
}




func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("os")
	address := r.FormValue("ip")
	host := r.FormValue("host")
    fmt.Fprintf(w, "os = %s\n", name)
	fmt.Fprintf(w, "ip = %s\n", address)
    fmt.Fprintf(w, "name = %s\n", host)
    EnterIntoDB(host, name, address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    /*
    if EntryExists("ip", r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")]) { 
        fmt.Fprintf(w, "ls")
    }
    */

    fmt.Fprintf(w, "ls ../")
    


}


func main() {
    if !isRoot()  {
        fmt.Println("please run as root...")
        os.Exit(1)
    }
    initializeMySQL()
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}