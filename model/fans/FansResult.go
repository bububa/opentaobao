package fans

// FansResult 结构体
type FansResult struct {
	// 推送成功列表
	DataList []bool `json:"data_list,omitempty" xml:"data_list>bool,omitempty"`
	// 失败message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 无意义
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 返回data
	Data *CashPoolVo `json:"data,omitempty" xml:"data,omitempty"`
	// 调用成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
