package main

import (
    "bytes"
    "database/sql"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // create and open datavese
    db, err := sql.Open("sqlite3", "./todo.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    // make sure connection is available
    err = db.Ping()
    if err != nil {
        fmt.Print(err.Error())
    }

    // create a struct to hold data
    type Task struct {
        Id         int
        Content    string
        CreatedAt  string
        UpdatedAt  string
    }
    router := gin.Default()

    // GET - get a task detail
    router.GET("/tasks/:id", func(c *gin.Context) {
        var (
            task Task
            result gin.H
        )
        id := c.Param("id")
        row := db.QueryRow("SELECT id, content, created_at, updated_at FROM tasks WHERE id = ?;", id)
        err = row.Scan(&task.Id, &task.Content, &task.CreatedAt, &task.UpdatedAt)
        if err != nil {
            // If no results send null
            result = gin.H{
                "result": nil,
                "count":  0,
            }
        } else {
            result = gin.H{
                "result": task,
                "count":  1,
            }
        }
        c.JSON(http.StatusOK, result)
    })

    // GET - retrieve all tasks
    router.GET("/tasks", func(c *gin.Context) {
        var (
            task  Task
            tasks []Task
        )
        rows, err := db.Query("SELECT id, content, created_at, updated_at FROM tasks;")
        if err != nil {
            fmt.Print(err.Error())
        }
        for rows.Next() {
            err = rows.Scan(&task.Id, &task.Content, &task.CreatedAt, &task.UpdatedAt)
            tasks = append(tasks, task)
            if err != nil {
                fmt.Print(err.Error())
            }
        }
        defer rows.Close()
        c.JSON(http.StatusOK, gin.H{
            "result": tasks,
            "count":  len(tasks),
        })
    })

    // POST - create task
    router.POST("/tasks", func(c *gin.Context) {
        var buffer bytes.Buffer
        content := c.PostForm("content")
        created_at := c.PostForm("created_at")
        updated_at := c.PostForm("updated_at")
        stmt, err := db.Prepare("INSERT INTO tasks (content, created_at, updated_at) values(?, ?, ?);")
        if err != nil {
            fmt.Print(err.Error())
        }
        _, err = stmt.Exec(content, created_at, updated_at)

        if err != nil {
            fmt.Print(err.Error())
        }

        // Fastest way to append strings
        buffer.WriteString(content)
        buffer.WriteString(" ")
        buffer.WriteString(created_at)
        buffer.WriteString(" ")
        buffer.WriteString(updated_at)
        defer stmt.Close()
        name := buffer.String()
        c.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf(" %s successfully created", name),
        })
    })

    // PUT - update task
    router.PUT("/tasks", func(c *gin.Context) {
        var buffer bytes.Buffer
        id := c.Query("id")
        content := c.PostForm("content")
        created_at := c.PostForm("created_at")
        updated_at := c.PostForm("updated_at")
        stmt, err := db.Prepare("UPDATE tasks SET content = ?, created_at = ?, updated_at = ? WHERE id = ?;")
        if err != nil {
            fmt.Print(err.Error())
        }
        _, err = stmt.Exec(content, created_at, updated_at, id)
        if err != nil {
            fmt.Print(err.Error())
        }

        // Fastest way to append strings
        buffer.WriteString(content)
        buffer.WriteString(" ")
        buffer.WriteString(created_at)
        buffer.WriteString(" ")
        buffer.WriteString(updated_at)
        defer stmt.Close()
        name := buffer.String()
        c.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("Successfully updated to %s", name),
        })
    })

    // DELETE - remove task
    router.DELETE("/tasks", func(c *gin.Context) {
        id := c.Query("id")
        stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?;")
        if err != nil {
            fmt.Print(err.Error())
        }
        _, err = stmt.Exec(id)
        if err != nil {
            fmt.Print(err.Error())
        }
        c.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("Successfully deleted user: %s", id),
        })
    })
    router.Run(":4000")
}