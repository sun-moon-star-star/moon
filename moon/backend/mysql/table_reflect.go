package mysql

import (
	"strings"
)

func GetTableType(table string)interface{} {
	table = strings.ToLower(table)

	if table == "trace_trace" {
		return &TraceTrace{}
	} else if table == "trace_span" {
		return &TraceSpan{}
	} else if table == "trace_span_reference" {
		return &TraceSpanReference{}
	} else if table == "trace_log" {
		return &TraceLog{}
	} else if table == "trace_tag" {
		return &TraceTag{}
	} else if table == "trace_baggage" {
		return &TraceBaggage{}
	} else if table == "log_log" {
		return &LogLog{}
	}

	return nil
}

func GetTableArrayType(table string) interface{} {
	table = strings.ToLower(table)

	if table == "trace_trace" {
		return &[]TraceTrace{}
	} else if table == "trace_span" {
		return &[]TraceSpan{}
	} else if table == "trace_span_reference" {
		return &[]TraceSpanReference{}
	} else if table == "trace_log" {
		return &[]TraceLog{}
	} else if table == "trace_tag" {
		return &[]TraceTag{}
	} else if table == "trace_baggage" {
		return &[]TraceBaggage{}
	} else if table == "log_log" {
		return &[]LogLog{}
	}

	return nil
}