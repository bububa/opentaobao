package alihouse

// AlibabaAlihouseExistinghomeBankcardSyncResult 结构体
type AlibabaAlihouseExistinghomeBankcardSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 唯一id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
