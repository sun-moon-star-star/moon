package http

import (
	"bytes"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
)

func TestConn(writer http.ResponseWriter, request *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(request.Body)
	fmt.Fprintf(writer, buf.String())
}

func SendHttp(writer http.ResponseWriter, request *http.Request) {
    // 根据请求body创建一个json解析器实例
    request.ParseForm()

    url := request.Form["url"][0]
    req := request.Form["req"][0]
    url = strings.Trim(url,"\"")

    resp, _ := http.Post(url, "application/json", bytes.NewBuffer([]byte(req)))
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(writer, string(body))
}