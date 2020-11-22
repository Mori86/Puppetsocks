package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func getTableResult(table *sql.Rows) string{ 
    defer table.Close()
    var ip string
    for table.Next() { 
        err := table.Scan(&ip)
	    if err != nil {
		    log.Fatal(err)
        }
        
	    
    }

    return (ip[len(ip)-1:])
}


func EntryExists(identifier string, value string) bool{
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/puppets")
    if err != nil {
        fmt.Println(err.Error()) 
    }

    defer db.Close()
    query := "SELECT EXISTS(SELECT * from set1 WHERE " + identifier + "=\"" + value + "\");"
    check, err := db.Query(query)
    
    
    if err != nil { 
        panic(err.Error())
    }
    defer check.Close()

    if getTableResult(check) == "1"{
        return true
    }
    return false
   
}


func EnterIntoDB(name string, os string, ip string) { 
   // fmt.Println("placing puppet into db...")

    if(EntryExists("ip", ip)) { 
        fmt.Println("ignoring duplicate host: ", ip, "...")
        return 
    }
   

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/puppets")
    if err != nil {
        fmt.Println(err.Error()) 
    }
    defer db.Close()
    query := "INSERT INTO set1 VALUES (\"" + name + "\", \"" + os + "\", \"" + ip + "\");"
    insert, err := db.Query(query)
    if err != nil { 
        panic(err.Error())
    }

    defer insert.Close()

    fmt.Println("[+] New connection from " + ip + "...")
}



func initializeMySQL() { 
    fmt.Println("connecting to database...")
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/puppets")
    if err != nil {
        fmt.Println(err.Error())
        
    }
    err = db.Ping()
    if err != nil {
        fmt.Println("database not connected...")
        panic(err)
    }
    defer db.Close()
    fmt.Println("Database is live..")

}
