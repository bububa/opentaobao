package omniorder

// TaobaoJstAstrolabeStoreinventoryItemupdateError 结构体
type TaobaoJstAstrolabeStoreinventoryItemupdateError struct {
	// 错误描述
	Descrpition string `json:"descrpition,omitempty" xml:"descrpition,omitempty"`
	// 处理失败的流水号
	FailedBillNum string `json:"failed_bill_num,omitempty" xml:"failed_bill_num,omitempty"`
}
