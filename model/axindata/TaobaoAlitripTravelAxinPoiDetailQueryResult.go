package axindata

// TaobaoalitriptravelaxinpoidetailqueryResult 结构体
type TaobaoalitriptravelaxinpoidetailqueryResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// poi详情
	Data *PoiDetaiVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
