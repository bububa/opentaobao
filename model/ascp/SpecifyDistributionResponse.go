package ascp

// SpecifyDistributionResponse 结构体
type SpecifyDistributionResponse struct {
	// 每个品的处理结果
	ResponseDetailList []SpecifyDistributionResponseDetail `json:"response_detail_list,omitempty" xml:"response_detail_list>specify_distribution_response_detail,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误消息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	//  处理是否成功，只有处理成功responseDetailList 中才会有值。 responseDetailList中代表每个品的处理结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
