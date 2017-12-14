### Build and Run migrate

`go build -o migrate.goapp migrate.go`

`./migrate.goapp`

#### or just

`go run migrate.go`

### Build and Run server

`go build -o todo.goapp todo.go`

`./todo.goapp`

#### or just

`go run todo.go`

### Curl tests

`curl -X POST --data "content=Creation of a new Task, test it;created_at=2017-12-12 12:00:00;updated_at=2017-12-12 12:00:00" localhost:4000/tasks`

`curl -X PUT --data "content=Update a Task Content;created_at=2017-12-12 12:00:00;updated_at=2017-12-12 13:00:00" localhost:4000/tasks?id=1`

`curl -X GET --data "" localhost:4000/tasks`

`curl -X DELETE --data "" localhost:4000/tasks?id=1`