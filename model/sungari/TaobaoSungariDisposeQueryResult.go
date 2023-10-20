package sungari

// TaobaosungaridisposequeryResult 结构体
type TaobaosungaridisposequeryResult struct {
	// data
	List []DisposeResultVo `json:"list,omitempty" xml:"list>dispose_result_vo,omitempty"`
	// 提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 服务是否调用成功，1：成功 2：失败 11：重复提交 其他：失败
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
