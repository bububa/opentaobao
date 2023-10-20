package baichuanctg

// CtgResponse 结构体
type CtgResponse struct {
	// 数据列表
	DataList []AlibabaBaichuanCtgToutiaoContentData `json:"data_list,omitempty" xml:"data_list>alibaba_baichuan_ctg_toutiao_content_data,omitempty"`
	// 返回错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否包含下一页数据
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 处理成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
