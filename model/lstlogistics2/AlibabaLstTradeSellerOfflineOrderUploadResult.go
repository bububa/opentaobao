package lstlogistics2

// AlibabaLstTradeSellerOfflineOrderUploadResult 结构体
type AlibabaLstTradeSellerOfflineOrderUploadResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 系统自动生成
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 出货单id
	Content int64 `json:"content,omitempty" xml:"content,omitempty"`
}
