package btrip

// OpenInvoiceModifyAndNewRq 结构体
type OpenInvoiceModifyAndNewRq struct {
	// 注册地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 银行账号
	BankNo string `json:"bank_no,omitempty" xml:"bank_no,omitempty"`
	// 开户行
	BankName string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 税号
	TaxNo string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	// 公司电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 第三方发票id
	ThirdPartId string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	// 发票抬头
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 类型，1:增值税普通发票,2:增值税专用发票
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
