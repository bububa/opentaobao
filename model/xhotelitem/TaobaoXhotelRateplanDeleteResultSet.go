package xhotelitem

// TaobaoxhotelrateplandeleteResultSet 结构体
type TaobaoxhotelrateplandeleteResultSet struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 房价id
	Rpid string `json:"rpid,omitempty" xml:"rpid,omitempty"`
}
