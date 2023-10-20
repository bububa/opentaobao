package alsc

import (
	"sync"
)

// TopInvoiceInsuranceNoDto 结构体
type TopInvoiceInsuranceNoDto struct {
	// 保险单号列表
	InsuranceNoList []string `json:"insurance_no_list,omitempty" xml:"insurance_no_list>string,omitempty"`
	// 发票申请单id
	InvoiceApplyId string `json:"invoice_apply_id,omitempty" xml:"invoice_apply_id,omitempty"`
	// 保险类型，医美险 MEDICAL_BEAUTY ，提前放款险 ADVANCE_LOAN
	InsuranceType string `json:"insurance_type,omitempty" xml:"insurance_type,omitempty"`
	// 统一社会信用代码
	UniqueCode string `json:"unique_code,omitempty" xml:"unique_code,omitempty"`
	// 月份，多个月份用逗号分隔
	Months string `json:"months,omitempty" xml:"months,omitempty"`
	// email：开电子发票 ，paper：开纸质发票
	PostType string `json:"post_type,omitempty" xml:"post_type,omitempty"`
	// 发票抬头
	OperatingLicense string `json:"operating_license,omitempty" xml:"operating_license,omitempty"`
	// 纳税人识别号
	TaxRegNumber string `json:"tax_reg_number,omitempty" xml:"tax_reg_number,omitempty"`
	// 发票地址
	InvoiceAddress string `json:"invoice_address,omitempty" xml:"invoice_address,omitempty"`
	// 发票电话
	InvoicePhone string `json:"invoice_phone,omitempty" xml:"invoice_phone,omitempty"`
	// 收件人地址
	InvoicePostAddress string `json:"invoice_post_address,omitempty" xml:"invoice_post_address,omitempty"`
	// 收件人号码
	InvoicePostPhone string `json:"invoice_post_phone,omitempty" xml:"invoice_post_phone,omitempty"`
	// 开户行
	BankName string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 开户行账号
	BankAccount string `json:"bank_account,omitempty" xml:"bank_account,omitempty"`
	// 电子邮件
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 发票金额 单位:分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 0 普通 1 专票
	InvoiceType int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
}

var poolTopInvoiceInsuranceNoDto = sync.Pool{
	New: func() any {
		return new(TopInvoiceInsuranceNoDto)
	},
}

// GetTopInvoiceInsuranceNoDto() 从对象池中获取TopInvoiceInsuranceNoDto
func GetTopInvoiceInsuranceNoDto() *TopInvoiceInsuranceNoDto {
	return poolTopInvoiceInsuranceNoDto.Get().(*TopInvoiceInsuranceNoDto)
}

// ReleaseTopInvoiceInsuranceNoDto 释放TopInvoiceInsuranceNoDto
func ReleaseTopInvoiceInsuranceNoDto(v *TopInvoiceInsuranceNoDto) {
	v.InsuranceNoList = v.InsuranceNoList[:0]
	v.InvoiceApplyId = ""
	v.InsuranceType = ""
	v.UniqueCode = ""
	v.Months = ""
	v.PostType = ""
	v.OperatingLicense = ""
	v.TaxRegNumber = ""
	v.InvoiceAddress = ""
	v.InvoicePhone = ""
	v.InvoicePostAddress = ""
	v.InvoicePostPhone = ""
	v.BankName = ""
	v.BankAccount = ""
	v.Email = ""
	v.Amount = 0
	v.InvoiceType = 0
	poolTopInvoiceInsuranceNoDto.Put(v)
}
