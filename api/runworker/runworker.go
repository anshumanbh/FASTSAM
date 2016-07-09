package main

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/errors"

	//	the below import statement should look like this - "github.com/<username>/FASTSAM/api/tasks" assuming
	//  you have the tasks.go file in github.com/<username>/FASTSAM/api/tasks folder
	//  I have my tasks.go file in github.com/bharta1/pentest-api-go/tasks folder hence the below import statement
	"github.com/bharta1/pentest-api-go/tasks"
)

var (
	cnf    config.Config
	server *machinery.Server
	worker *machinery.Worker
)

func init() {
	cnf = config.Config{
		Broker:        "amqp://guest:guest@localhost:5672/",
		ResultBackend: "amqp://guest:guest@localhost:5672/",
		Exchange:      "machinery_exchange",
		ExchangeType:  "direct",
		DefaultQueue:  "machinery_tasks",
		BindingKey:    "machinery_task",
	}

	server, err := machinery.NewServer(&cnf)
	errors.Fail(err, "Could not initialize server")

	// Register tasks
	tasks := map[string]interface{}{
		"NmapScan": tasks.NmapScan,
		"MasScan":  tasks.MasScan,
	}
	server.RegisterTasks(tasks)

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker = server.NewWorker("machinery_worker")
}

func main() {
	err := worker.Launch()
	errors.Fail(err, "Could not launch worker")
}
