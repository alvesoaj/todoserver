
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "os"
)

func main() {
    // remove database
    os.Remove("./todo.db")

    // create and open datavese
    db, err := sql.Open("sqlite3", "./todo.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    // make sure connection is available
    err = db.Ping()
    if err != nil {
        fmt.Println(err.Error())
    }

    // create tasks table
    stmt, err := db.Prepare("CREATE TABLE tasks (id INTEGER PRIMARY KEY, content TEXT, created_at DATETIME, updated_at DATETIME);")
    if err != nil {
        fmt.Println(err.Error())
    }
    _, err = stmt.Exec()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Tasks table successfully migrated....")
    }
}