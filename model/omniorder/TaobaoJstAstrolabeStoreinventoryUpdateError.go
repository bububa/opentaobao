package omniorder

// TaobaojstastrolabestoreinventoryupdateError 结构体
type TaobaojstastrolabestoreinventoryupdateError struct {
	// 错误描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 处理失败的流水号
	FailedBillNum string `json:"failed_bill_num,omitempty" xml:"failed_bill_num,omitempty"`
}
