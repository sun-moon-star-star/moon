package mysql

import (
	"strings"
)

func GetTableType(table string)interface{} {
	table = strings.ToLower(table)

	if table == "trace" {
		return &trace{}
	} else if table == "span" {
		return &span{}
	} else if table == "span_reference" {
		return &span_reference{}
	} else if table == "log" {
		return &log{}
	} else if table == "tag" {
		return &tag{}
	} else if table == "baggage" {
		return &baggage{}
	}

	return nil
}

func GetTableArrayType(table string) interface{} {
	table = strings.ToLower(table)

	if table == "trace" {
		return &[]trace{}
	} else if table == "span" {
		return &[]span{}
	} else if table == "span_reference" {
		return &[]span_reference{}
	} else if table == "log" {
		return &[]log{}
	} else if table == "tag" {
		return &[]tag{}
	} else if table == "baggage" {
		return &[]baggage{}
	}

	return nil
}