package alihouse

// AlibabaAlihouseNewhomeDepositPublishResult 结构体
type AlibabaAlihouseNewhomeDepositPublishResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 预存金淘宝商品id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
