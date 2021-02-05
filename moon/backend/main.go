package main

import (
    "log"
    "net/http"
    main_http "main/http"
    main_mysql "main/mysql"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/")))
	http.HandleFunc("/testConn", main_http.TestConn)
    http.HandleFunc("/sendHttp", main_http.SendHttp)
    http.HandleFunc("/loadTables", main_mysql.LoadTables)
    http.HandleFunc("/selectTableLimit", main_mysql.SelectTableLimit)
    err := http.ListenAndServe(":8151", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
