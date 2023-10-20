package dutyfree

// AlibabadutyfreestockqueryResultDo 结构体
type AlibabadutyfreestockqueryResultDo struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 具体库存信息
	Object *StockResultDto `json:"object,omitempty" xml:"object,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
