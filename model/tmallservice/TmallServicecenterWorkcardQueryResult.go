package tmallservice

// TmallServicecenterWorkcardQueryResult 结构体
type TmallServicecenterWorkcardQueryResult struct {
	// 分页数据
	ResultData *Paged `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
