package main

import (
    "log"
    "net/http"
    main_http "main/http"
    main_mysql "main/mysql"
    main_text "main/text"
)

func main() {
    main_text.RandSeedTime()
	http.Handle("/", http.FileServer(http.Dir("/")))
	http.HandleFunc("/testConn", main_http.TestConn)
    http.HandleFunc("/sendHttp", main_http.SendHttp)
    http.HandleFunc("/loadTables", main_mysql.LoadTables)
    http.HandleFunc("/selectTableLimit", main_mysql.SelectTableLimit)
    http.HandleFunc("/updateTableField", main_mysql.UpdateTableField)
    http.HandleFunc("/insertTable", main_mysql.InsertTable)
    http.HandleFunc("/savePaste", main_text.SavePaste)
    http.HandleFunc("/loadPaste", main_text.LoadPaste)
    err := http.ListenAndServe(":8151", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
