package mysql

import (
	"strings"
)

func GetTableType(table string)interface{} {
	table = strings.ToLower(table)

	if table == "trace" {
		return &Trace{}
	} else if table == "span" {
		return &Span{}
	} else if table == "span_reference" {
		return &SpanReference{}
	} else if table == "log" {
		return &Log{}
	} else if table == "tag" {
		return &Tag{}
	} else if table == "baggage" {
		return &Baggage{}
	}

	return nil
}

func GetTableArrayType(table string) interface{} {
	table = strings.ToLower(table)

	if table == "trace" {
		return &[]Trace{}
	} else if table == "span" {
		return &[]Span{}
	} else if table == "span_reference" {
		return &[]SpanReference{}
	} else if table == "log" {
		return &[]Log{}
	} else if table == "tag" {
		return &[]Tag{}
	} else if table == "baggage" {
		return &[]Baggage{}
	}

	return nil
}