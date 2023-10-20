package singletreasure

// TaobaosingletreasureactivityitembatchupdateResultDto 结构体
type TaobaosingletreasureactivityitembatchupdateResultDto struct {
	// 返回的描述信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回所有的处理错误的信息
	Data *ItemProcessErrorResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 处理结果
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
