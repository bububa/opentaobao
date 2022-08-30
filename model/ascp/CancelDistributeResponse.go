package ascp

// CancelDistributeResponse 结构体
type CancelDistributeResponse struct {
	// 处理结果。 打平到 item + 分销商粒度。 处理成功和处理失败都会返回
	CancelDistributeDetailList []CancelDistributeDetail `json:"cancel_distribute_detail_list,omitempty" xml:"cancel_distribute_detail_list>cancel_distribute_detail,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误内容
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
