package xhotelitem

// TaobaoxhotelitemselectionsellerstatexposureResult 结构体
type TaobaoxhotelitemselectionsellerstatexposureResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Module *TaobaoxhotelitemselectionsellerstatexposureModule `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
