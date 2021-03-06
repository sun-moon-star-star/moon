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

func getConn(params map[string]interface{})(*gorm.DB, error) {
	hostname := params["hostname"]
	port := params["port"]
	username := params["username"]
	password := params["password"]
	network := params["network"]
	database := params["database"]
	
	db_desc := fmt.Sprintf("%v:%v@%v(%v:%v)/%v",username,password,network,hostname,port,database)
	return getConnByDesc(db_desc)
}

func LoadTables(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]interface{}
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
	var params map[string]interface{}
	decoder.Decode(&params)

	db, err := getConn(params)
	
	ret := make(map[string]interface{})

	if (err != nil) {
		goto err
	} else {
		database := params["database"].(string)
		table := params["table"].(string)
		num := params["num"]

		data := GetTableArrayType(database + "_" + table)
		if data == nil {
			goto err
		}
		res := db.Table(table).Find(data).Limit(num)

		if res.Error != nil {
			err = res.Error
			goto err
		}
	
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

func UpdateTableField(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]interface{}
	decoder.Decode(&params)

	db, err := getConn(params)
	
	ret := make(map[string]interface{})

	if (err != nil) {
		goto err
	} else {
		table := params["table"].(string)
		unique_field := params["unique_field"].(string)
		unique_field_value := params["unique_field_value"]
		field := params["field"]
		value := params["value"]

		data := GetTableType(table)
		if data == nil {
			goto err
		}

		res := db.Table(table).Where(unique_field + " = ?", unique_field_value).Update(field, value)
		if res.Error != nil {
			err = res.Error
			goto err
		}
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

func InsertTable(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]interface{}
	decoder.Decode(&params)

	db, err := getConn(params)
	
	ret := make(map[string]interface{})

	if (err != nil) {
		goto err
	} else {
		database := params["database"].(string)
		table := params["table"].(string)
		object_jsonstr := params["object_jsonstr"].(string) // json string

		data := GetTableType(database + "_" + table)
		if data == nil {
			goto err
		}

		err = json.Unmarshal([]byte(object_jsonstr), data)

		res := db.Table(table).Create(data)

		if res.Error != nil {
			err = res.Error
			goto err
		}
	
		ret["data"] = res
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