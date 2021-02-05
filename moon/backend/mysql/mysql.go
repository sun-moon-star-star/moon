package mysql

import (
    "net/http"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"time"
)

var pool map[string]*gorm.DB

func getConnByDesc(db_desc string)(*gorm.DB, error) {
	if pool == nil {
		pool = make(map[string]*gorm.DB)
	}

	if db, ok := pool[db_desc]; ok {
		return db, nil
	}

	db, err := gorm.Open("mysql", db_desc)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	
	db.DB().SetConnMaxLifetime(time.Duration(3600) * time.Second)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)

	pool[db_desc] = db
	return db, nil
}

func getConn(params map[string]string)(*gorm.DB, error) {
	hostname := params["hostname"]
	port := params["port"]
	username := params["username"]
	password := params["password"]
	network := params["network"]
	database := params["database"]
	
	db_desc := fmt.Sprintf("%v:%v@%v(%v:%v)/%v",username,password,network,hostname,port,database)
	fmt.Println(db_desc)
	return getConnByDesc(db_desc)
}

func LoadTables(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]string
	decoder.Decode(&params)

	db, err := getConn(params)
	
	ret := make(map[string]interface{})

	if err != nil {
		goto err
	} else {
		res, err := db.DB().Query("SHOW TABLES")
		if err != nil {
			goto err
		}
		defer res.Close()
	
		var table string
		var tables []string

		for res.Next() {
			res.Scan(&table)
			tables = append(tables, table)
		}

		ret["data"] = tables
		goto end
	}

err:
	ret["error"] = -1
	ret["msg"] = err.Error()
	goto ret
end:
	ret["error"] = 0
	ret["msg"] = "success"
ret:

	var ret_bytes []byte
	ret_bytes, err = json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}
	fmt.Fprintf(writer, string(ret_bytes))
}

func SelectTableLimit(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]string
	decoder.Decode(&params)

	db, err := getConn(params)
	
	ret := make(map[string]interface{})

	if (err != nil) {
		goto err
	} else {
		table := params["table"]
		num := params["num"]

		data := GetTableArrayType(table)
		if data == nil {
			goto err
		}
		db.Table(table).Find(data).Limit(num)
	
		ret["data"] = data
		goto end
	}

err:
	ret["error"] = -1
	ret["msg"] = err.Error()
	goto ret
end:
	ret["error"] = 0
	ret["msg"] = "success"
ret:

	var ret_bytes []byte
	ret_bytes, err = json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}
	fmt.Fprintf(writer, string(ret_bytes))
}

