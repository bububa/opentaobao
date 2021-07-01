package cainiaolocker

// SingleResult 结构体
type SingleResult struct {
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 参照返回码定义列表
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// data
	Data *CainiaoEndpointLockerTopOrderNoticesendQueryData `json:"data,omitempty" xml:"data,omitempty"`
}
