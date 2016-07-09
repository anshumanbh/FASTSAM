package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/errors"
	"github.com/RichardKnop/machinery/v1/signatures"
)

var (
	server *machinery.Server

	cnf = config.Config{
		Broker:        "amqp://guest:guest@localhost:5672/",
		ResultBackend: "amqp://guest:guest@localhost:5672/",
		Exchange:      "machinery_exchange",
		ExchangeType:  "direct",
		DefaultQueue:  "machinery_tasks",
		BindingKey:    "machinery_task",
	}

	path = "/tmp/iplist.txt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func NmapScan(w http.ResponseWriter, r *http.Request) {

	var ips []string

	body, err := ioutil.ReadAll(r.Body)
	check(err)

	err = r.Body.Close()
	check(err)

	err = json.Unmarshal(body, &ips)
	check(err)

	var buffer bytes.Buffer

	buffer.WriteString("TARGETS=")

	for _, element := range ips {
		buffer.WriteString(element)
		buffer.WriteString(" ")
	}

	tstring := buffer.String()

	task := signatures.TaskSignature{
		Name: "NmapScan",
		Args: []signatures.TaskArg{
			{
				Type:  "string",
				Value: tstring,
			},
		},
	}

	server, err := machinery.NewServer(&cnf)
	errors.Fail(err, "Could not initialize server")

	asyncResult, err := server.SendTask(&task)
	check(err)

	taskstate := asyncResult.GetState()

	w.Write([]byte("NmapScan job was submitted. The task ID is " + taskstate.TaskUUID + "\n"))

}

func MasScan(w http.ResponseWriter, r *http.Request) {

	var ips []string

	body, err := ioutil.ReadAll(r.Body)
	check(err)

	err = r.Body.Close()
	check(err)

	err = json.Unmarshal(body, &ips)
	check(err)

	var buffer bytes.Buffer

	buffer.WriteString("TARGETS=")

	for _, element := range ips {
		buffer.WriteString(element)
		buffer.WriteString(" ")
	}

	tstring := buffer.String()

	task := signatures.TaskSignature{
		Name: "MasScan",
		Args: []signatures.TaskArg{
			{
				Type:  "string",
				Value: tstring,
			},
		},
	}

	server, err := machinery.NewServer(&cnf)
	errors.Fail(err, "Could not initialize server")

	asyncResult, err := server.SendTask(&task)
	check(err)

	taskstate := asyncResult.GetState()

	w.Write([]byte("MasScan job was submitted. The task ID is " + taskstate.TaskUUID + "\n"))

}
