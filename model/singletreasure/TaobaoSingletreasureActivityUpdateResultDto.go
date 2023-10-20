package singletreasure

// TaobaosingletreasureactivityupdateResultDto 结构体
type TaobaosingletreasureactivityupdateResultDto struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 套餐编辑是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 系统是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
