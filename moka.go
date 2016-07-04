package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var metadataDefinitions string
var projectMetadataDefinition string
var ingestMetadataDefinition string
var publishMetadataDefinition string

var workflowDefinitions string
var runWorkflow string
var workflowStatus string

var loginCreds string

func SetHeaders(w http.ResponseWriter) (http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Headers", "requested-with, Content-Type, origin, authorization, accept, client-security-token")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Max-Age", "10000")

    w.Header().Set("Content-Type", "application/json")

    return w
}

func MetadataByID(c web.C, w http.ResponseWriter, r *http.Request) {
    // A separate method for this, outside of moka.Router, is pure laziness
    // - it was easier than building regexp matchers

    w = SetHeaders(w)

    switch c.URLParams["id"] {
    case "1":
        fmt.Fprintf(w, projectMetadataDefinition)
    case "2":
        fmt.Fprintf(w, ingestMetadataDefinition)
    case "3":
        fmt.Fprintf(w, publishMetadataDefinition)
    }
}

func Status() (string) {
    name, _ := os.Hostname()
    return fmt.Sprintf("{\"status\":\"Healthy!\",\"host\":\"%s\"}", name)
}

func Router(c web.C, w http.ResponseWriter, r *http.Request) {
    w = SetHeaders(w)

    switch r.URL.Path {
    case "/metadataDefinitions":
        fmt.Fprintf(w, metadataDefinitions)
    case "/workflowDefinitions":
        fmt.Fprintf(w, workflowDefinitions)
    case "/workflows":
        fmt.Fprintf(w, runWorkflow)
    case "/login":
        fmt.Fprintf(w, loginCreds)
    case "/workflows/1","/workflows/2", "/workflows/3":
        fmt.Fprintf(w, workflowStatus)

    default:
        fmt.Fprintf(w, Status())
    }
}

func LoadJson(f string) (string){
    output,err := ioutil.ReadFile(f)
    if err != nil {
        log.Fatal(err)
    }

    return string(output)
}

func main() {
    metadataDefinitions = LoadJson("./json/metadataDefinitions.json")
    projectMetadataDefinition = LoadJson("./json/projectMetadataDefinition.json")
    ingestMetadataDefinition  = LoadJson("./json/ingestMetadataDefinition.json")
    publishMetadataDefinition = LoadJson("./json/publishMetadataDefinition.json")

    workflowDefinitions = LoadJson("./json/workflowDefinitions.json")
    runWorkflow = LoadJson("./json/runWorkflow.json")

    loginCreds = LoadJson("./json/loginCreds.json")

    goji.Get("/", Router)
    goji.Options("/*", Router)

    goji.Get("/metadataDefinitions", Router)
    goji.Get("/metadataDefinitions/:id/definition", MetadataByID)

    goji.Get("/workflowDefinitions", Router)

    goji.Post("/workflows", Router)
    goji.Get("/workflows/:id", Router)

    goji.Post("/login", Router)

    goji.Serve()
}
