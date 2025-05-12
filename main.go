package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Customers struct {
    XMLName   xml.Name   `xml:"Customers"`
    Customers []Customer `xml:"Customer"`
}

type Customer struct {
    Document   string `xml:"Document"`
    Name       string `xml:"Name"`
    Address    string `xml:"Address"`
    Profession string `xml:"Profession"`
}

func main() {
    // Read XML file
    data, err := ioutil.ReadFile("customers.xml")
    if err != nil {
        panic(err)
    }

    var customers Customers
    err = xml.Unmarshal(data, &customers)
    if err != nil {
        panic(err)
    }

    // Connect to MySQL
    db, err := sql.Open("mysql", "appuser:apppassword@tcp(192.168.1.100:3306)/demo")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    for _, c := range customers.Customers {
        _, err := db.Exec("INSERT INTO customers (document, name, address, profession) VALUES (?, ?, ?, ?)",
            c.Document, c.Name, c.Address, c.Profession)
        if err != nil {
            fmt.Println("Error inserting:", err)
        } else {
            fmt.Println("Inserted:", c.Name)
        }
    }
}
