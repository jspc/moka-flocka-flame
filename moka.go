package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var metadataDefinitions []byte
var projectMetadataDefinition []byte
var ingestMetadataDefinition []byte
var publishMetadataDefinition []byte

var workflowDefinitions []byte
var runWorkflow []byte
var workflowStatus []byte

func get_root(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Headers", "requested-with, Content-Type, origin, authorization, accept, client-security-token")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Max-Age", "10000")

    w.Header().Set("Content-Type", "application/json")

    name, _ := os.Hostname()
    fmt.Fprintf(w, "{\"status\":\"Healthy!\",\"host\":\"%s\"}", name)
}

func get_metadataDefinitions(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(metadataDefinitions))
}

func get_metadataDefinitionID(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var output []byte

    switch c.URLParams["id"] {
    case "1": output = projectMetadataDefinition
    case "2": output = ingestMetadataDefinition
    case "3": output = publishMetadataDefinition
    }

    fmt.Fprintf(w, string(output))
}

func get_workflowDefinitions(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(workflowDefinitions))
}

func post_workflows(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(runWorkflow))
}

func get_workflow_status(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(workflowStatus))
}


func main() {
    metadataDefinitions,_ = ioutil.ReadFile("./json/metadataDefinitions.json")
    projectMetadataDefinition ,_  = ioutil.ReadFile("./json/projectMetadataDefinition.json")
    ingestMetadataDefinition ,_   = ioutil.ReadFile("./json/ingestMetadataDefinition.json")
    publishMetadataDefinition ,_  = ioutil.ReadFile("./json/publishMetadataDefinition.json")

    workflowDefinitions,_ = ioutil.ReadFile("./json/workflowDefinitions.json")
    runWorkflow,_ = ioutil.ReadFile("./json/runWorkflow.json")


    goji.Get("/", get_root)
    goji.Options("/*", get_root)

    goji.Get("/metadataDefinitions", get_metadataDefinitions)
    goji.Get("/metadataDefinitions/:id/definition", get_metadataDefinitionID)

    goji.Get("/workflowDefinitions", get_workflowDefinitions)

    goji.Post("/workflows", post_workflows)
    goji.Get("/workflows/:id", get_workflow_status)

    goji.Serve()
}
