package fpm

// InvoiceMainExt 结构体
type InvoiceMainExt struct {
	// 申请人
	AuthApplyUserId string `json:"auth_apply_user_id,omitempty" xml:"auth_apply_user_id,omitempty"`
	// 扫描时间
	ScanTime string `json:"scan_time,omitempty" xml:"scan_time,omitempty"`
	// 销方名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 购方编码
	PurchaserCode string `json:"purchaser_code,omitempty" xml:"purchaser_code,omitempty"`
	// 发票密文（字符200位）
	CipherText string `json:"cipher_text,omitempty" xml:"cipher_text,omitempty"`
	// 申请时间
	AuthApplyTime string `json:"auth_apply_time,omitempty" xml:"auth_apply_time,omitempty"`
	// 发票类型
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 代开税号
	IssuedTaxNo string `json:"issued_tax_no,omitempty" xml:"issued_tax_no,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 附件的resourceId(上传影响接口返回)
	FileId string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// 不含税金额
	AmountWithoutTax string `json:"amount_without_tax,omitempty" xml:"amount_without_tax,omitempty"`
	// 批次号
	BatchNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// 机器编号（字符20位）
	MachineCode string `json:"machine_code,omitempty" xml:"machine_code,omitempty"`
	// 扫描员工号
	ScanUserId string `json:"scan_user_id,omitempty" xml:"scan_user_id,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 二维码发票标记
	TwoCodeFlag string `json:"two_code_flag,omitempty" xml:"two_code_flag,omitempty"`
	// 扫描账号
	ScanAccount string `json:"scan_account,omitempty" xml:"scan_account,omitempty"`
	// 购方名称
	PurchaserName string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 销方税号
	SellerTaxNo string `json:"seller_tax_no,omitempty" xml:"seller_tax_no,omitempty"`
	// 二维码密文 （字符1000位）	二维码发票必填
	CipherTextTwocode string `json:"cipher_text_twocode,omitempty" xml:"cipher_text_twocode,omitempty"`
	// 购方税号
	PurchaserTaxNo string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
	// 校验码
	VerifyNo string `json:"verify_no,omitempty" xml:"verify_no,omitempty"`
	// 发票开票日期
	PaperDrewDate string `json:"paper_drew_date,omitempty" xml:"paper_drew_date,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 快递号
	WaybillNo string `json:"waybill_no,omitempty" xml:"waybill_no,omitempty"`
}
