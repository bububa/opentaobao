package itpolicy

// ErrorFareRow 结构体
type ErrorFareRow struct {
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误行数
	RowNum int64 `json:"row_num,omitempty" xml:"row_num,omitempty"`
}
