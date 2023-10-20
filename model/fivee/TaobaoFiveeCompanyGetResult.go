package fivee

// TaobaofiveecompanygetResult 结构体
type TaobaofiveecompanygetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *Company `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
