package westcrm

// DataResult 结构体
type DataResult struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	DataList *CustomerBaseInfoVo `json:"data_list,omitempty" xml:"data_list,omitempty"`
	// 参数code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *UserConsumerVo `json:"data,omitempty" xml:"data,omitempty"`
}
