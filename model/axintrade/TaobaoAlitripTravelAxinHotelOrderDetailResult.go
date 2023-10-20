package axintrade

// TaobaoalitriptravelaxinhotelorderdetailResult 结构体
type TaobaoalitriptravelaxinhotelorderdetailResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回参数
	Data *HotelOrderQueryRes `json:"data,omitempty" xml:"data,omitempty"`
	// 成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
