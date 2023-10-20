package fpm

import (
	"sync"
)

// RegisterInvoiceDto 结构体
type RegisterInvoiceDto struct {
	// 发票行列表
	LineList []RegisterInvoiceLineDto `json:"line_list,omitempty" xml:"line_list>register_invoice_line_dto,omitempty"`
	// 发票备注
	InvoiceRemark string `json:"invoice_remark,omitempty" xml:"invoice_remark,omitempty"`
	// 销方国家代码
	SellerCountryCode string `json:"seller_country_code,omitempty" xml:"seller_country_code,omitempty"`
	// 购方银行信息
	PurchaserBankInfo string `json:"purchaser_bank_info,omitempty" xml:"purchaser_bank_info,omitempty"`
	// 销方银行信息
	SellerBankInfo string `json:"seller_bank_info,omitempty" xml:"seller_bank_info,omitempty"`
	// 购方地区编码
	PurchaserRegionCode string `json:"purchaser_region_code,omitempty" xml:"purchaser_region_code,omitempty"`
	// 购方公司代码
	PurchaserCode string `json:"purchaser_code,omitempty" xml:"purchaser_code,omitempty"`
	// 发票密文（字符200位）
	CipherText string `json:"cipher_text,omitempty" xml:"cipher_text,omitempty"`
	// 二维码密文
	QrCodeCipherText string `json:"qr_code_cipher_text,omitempty" xml:"qr_code_cipher_text,omitempty"`
	// 发票类型
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 文件下载http地址（优先级高于fileContent）
	FileDownloadHttpUrl string `json:"file_download_http_url,omitempty" xml:"file_download_http_url,omitempty"`
	// 机器编码
	MachineCode string `json:"machine_code,omitempty" xml:"machine_code,omitempty"`
	// 影像id
	ImageId string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 购方名称
	PurchaserName string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	// 扫描账号
	ScanAccount string `json:"scan_account,omitempty" xml:"scan_account,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 电子发票类型
	ElectronicType string `json:"electronic_type,omitempty" xml:"electronic_type,omitempty"`
	// 销方税号
	SellerTaxNo string `json:"seller_tax_no,omitempty" xml:"seller_tax_no,omitempty"`
	// 开票国家类型
	IssueCountryType string `json:"issue_country_type,omitempty" xml:"issue_country_type,omitempty"`
	// 登记成功时间
	RegTime string `json:"reg_time,omitempty" xml:"reg_time,omitempty"`
	// 关联单据号
	RelatedOrder string `json:"related_order,omitempty" xml:"related_order,omitempty"`
	// 购方税号
	PurchaserTaxNo string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
	// 影像原始文件名称
	ImageFileName string `json:"image_file_name,omitempty" xml:"image_file_name,omitempty"`
	// 登记渠道
	RegisterChannel string `json:"register_channel,omitempty" xml:"register_channel,omitempty"`
	// 校验码
	CheckSum string `json:"check_sum,omitempty" xml:"check_sum,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 代开销方名称
	IssuedSellerName string `json:"issued_seller_name,omitempty" xml:"issued_seller_name,omitempty"`
	// 不含税金额
	ExcludingTaxAmount string `json:"excluding_tax_amount,omitempty" xml:"excluding_tax_amount,omitempty"`
	// 所属平台
	RegPlatformCode string `json:"reg_platform_code,omitempty" xml:"reg_platform_code,omitempty"`
	// 二维码标识
	QrCodeFlag string `json:"qr_code_flag,omitempty" xml:"qr_code_flag,omitempty"`
	// 登记的OU
	RegBizIdentityOu string `json:"reg_biz_identity_ou,omitempty" xml:"reg_biz_identity_ou,omitempty"`
	// 销方名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 登记人
	RegUser string `json:"reg_user,omitempty" xml:"reg_user,omitempty"`
	// 销方code
	SellerCode string `json:"seller_code,omitempty" xml:"seller_code,omitempty"`
	// 购方国家代码
	PurchaserCountryCode string `json:"purchaser_country_code,omitempty" xml:"purchaser_country_code,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 代开单位税号
	IssuedTaxNo string `json:"issued_tax_no,omitempty" xml:"issued_tax_no,omitempty"`
	// 有效税额
	EffectiveTaxAmount string `json:"effective_tax_amount,omitempty" xml:"effective_tax_amount,omitempty"`
	// 购方地址及电话
	PurchaserContactInfo string `json:"purchaser_contact_info,omitempty" xml:"purchaser_contact_info,omitempty"`
	// 发票总金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 销方地区编码
	SellerRegionCode string `json:"seller_region_code,omitempty" xml:"seller_region_code,omitempty"`
	// 开票日期
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// 发票介质类型
	InvoiceMaterial string `json:"invoice_material,omitempty" xml:"invoice_material,omitempty"`
	// 销方地址及电话
	SellerContactInfo string `json:"seller_contact_info,omitempty" xml:"seller_contact_info,omitempty"`
	// 是否海外某些国家的发票使用
	AuthorizedDealer string `json:"authorized_dealer,omitempty" xml:"authorized_dealer,omitempty"`
	// 登记时，所属的业务身份
	RegBizIdentityCode string `json:"reg_biz_identity_code,omitempty" xml:"reg_biz_identity_code,omitempty"`
	// 运单号
	WaybillNo string `json:"waybill_no,omitempty" xml:"waybill_no,omitempty"`
	// 票种类型code
	InvoiceTypeCode string `json:"invoice_type_code,omitempty" xml:"invoice_type_code,omitempty"`
	// 登记时，最终所选的bizId
	RegBizIdentityId int64 `json:"reg_biz_identity_id,omitempty" xml:"reg_biz_identity_id,omitempty"`
	// 纸票是否到票才认证
	PaperOwnCertifyFlag bool `json:"paper_own_certify_flag,omitempty" xml:"paper_own_certify_flag,omitempty"`
}

var poolRegisterInvoiceDto = sync.Pool{
	New: func() any {
		return new(RegisterInvoiceDto)
	},
}

// GetRegisterInvoiceDto() 从对象池中获取RegisterInvoiceDto
func GetRegisterInvoiceDto() *RegisterInvoiceDto {
	return poolRegisterInvoiceDto.Get().(*RegisterInvoiceDto)
}

// ReleaseRegisterInvoiceDto 释放RegisterInvoiceDto
func ReleaseRegisterInvoiceDto(v *RegisterInvoiceDto) {
	v.LineList = v.LineList[:0]
	v.InvoiceRemark = ""
	v.SellerCountryCode = ""
	v.PurchaserBankInfo = ""
	v.SellerBankInfo = ""
	v.PurchaserRegionCode = ""
	v.PurchaserCode = ""
	v.CipherText = ""
	v.QrCodeCipherText = ""
	v.InvoiceType = ""
	v.InvoiceNo = ""
	v.FileDownloadHttpUrl = ""
	v.MachineCode = ""
	v.ImageId = ""
	v.InvoiceCode = ""
	v.PurchaserName = ""
	v.ScanAccount = ""
	v.TaxRate = ""
	v.ElectronicType = ""
	v.SellerTaxNo = ""
	v.IssueCountryType = ""
	v.RegTime = ""
	v.RelatedOrder = ""
	v.PurchaserTaxNo = ""
	v.ImageFileName = ""
	v.RegisterChannel = ""
	v.CheckSum = ""
	v.TaxAmount = ""
	v.IssuedSellerName = ""
	v.ExcludingTaxAmount = ""
	v.RegPlatformCode = ""
	v.QrCodeFlag = ""
	v.RegBizIdentityOu = ""
	v.SellerName = ""
	v.RegUser = ""
	v.SellerCode = ""
	v.PurchaserCountryCode = ""
	v.Currency = ""
	v.IssuedTaxNo = ""
	v.EffectiveTaxAmount = ""
	v.PurchaserContactInfo = ""
	v.Amount = ""
	v.SellerRegionCode = ""
	v.InvoiceDate = ""
	v.InvoiceMaterial = ""
	v.SellerContactInfo = ""
	v.AuthorizedDealer = ""
	v.RegBizIdentityCode = ""
	v.WaybillNo = ""
	v.InvoiceTypeCode = ""
	v.RegBizIdentityId = 0
	v.PaperOwnCertifyFlag = false
	poolRegisterInvoiceDto.Put(v)
}
