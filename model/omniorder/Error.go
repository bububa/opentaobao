package omniorder

// Error 
type Error struct {
    // 错误描述
    Descrpition   string `json:"descrpition,omitempty" xml:"descrpition,omitempty"`
    // 处理失败的流水号（有多个时，用逗号分隔）
    FailedBillNum   string `json:"failed_bill_num,omitempty" xml:"failed_bill_num,omitempty"`
    // 错误描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
}
