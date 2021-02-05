package text

import (
    "net/http"
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func SavePaste(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]interface{}
	decoder.Decode(&params)

	var err error

	ret := make(map[string]interface{})

	var text string
	var token string
	var filename string
	var file *os.File
	var nbytes int
	
	text = params["text"].(string)
	if token_obj, ok := params["token"]; !ok {
		token, err = GetUUID()
		if err != nil {
			goto err
		}	
	} else {
		token = token_obj.(string)
	}

	filename = "../data/paste/" + token + ".txt"

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		file, err = os.Create(filename)
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY, 0777)
	}

	if err != nil {
		goto err
	}	

	defer file.Close()
	nbytes, err = io.WriteString(file, text)
	if err != nil {
		goto err
	}
	
	ret["error"] = 0
	ret["msg"] = "success"
	ret["token"] = token
	ret["bytes"] = nbytes
	goto ret
err:
	ret["error"] = -1
	ret["msg"] = err.Error()
ret:

	var ret_bytes []byte
	ret_bytes, err = json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}
	fmt.Fprintf(writer, string(ret_bytes))
}


func LoadPaste(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]interface{}
	decoder.Decode(&params)

	var err error

	ret := make(map[string]interface{})

	var token string
	var filename string
	var file *os.File
	var nbytes int
	var bytes []byte
	
	if token_obj, ok := params["token"]; !ok {
		err = fmt.Errorf("file not set")
		goto err
	} else {
		token = token_obj.(string)
	}

	filename = "../data/paste/" + token + ".txt"

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		err = fmt.Errorf("file not exist: %v", filename)
		goto err
	} else {
		file, err = os.OpenFile(filename, os.O_RDONLY, 0777)
	}

	if err != nil {
		goto err
	}	

	defer file.Close()
	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		goto err
	}
	
	ret["error"] = 0
	ret["msg"] = "success"
	ret["token"] = token
	ret["text"] = string(bytes)
	ret["bytes"] = nbytes
	goto ret
err:
	ret["error"] = -1
	ret["msg"] = err.Error()
ret:

	var ret_bytes []byte
	ret_bytes, err = json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}
	fmt.Fprintf(writer, string(ret_bytes))
}