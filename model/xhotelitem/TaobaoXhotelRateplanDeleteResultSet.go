package xhotelitem

// TaobaoXhotelRateplanDeleteResultSet 结构体
type TaobaoXhotelRateplanDeleteResultSet struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 房价id
	Rpid string `json:"rpid,omitempty" xml:"rpid,omitempty"`
}
