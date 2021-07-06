package inventory

// ResultCode 结构体
type ResultCode struct {
	// 结果描述
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果码
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
