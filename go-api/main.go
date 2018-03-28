package main

import (
  "net/http"
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
  "go-api/Cassandra"
  "go-api/Licences"
)

type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

func main() {
  CassandraSession := Cassandra.Session
    defer CassandraSession.Close()

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", heartbeat)
  router.HandleFunc("/licence/{licence_LicenceRSN}", Licences.GetLicence)
  log.Fatal(http.ListenAndServe(":8081", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
