package jst

// TaobaordsdbcreateaccountResultSet 结构体
type TaobaordsdbcreateaccountResultSet struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// exception
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// totalResults
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
