package xhotelonlineorder

import (
	"sync"
)

// CommonInvoiceInfo 结构体
type CommonInvoiceInfo struct {
	// 发票抬头
	CompanyTitle string `json:"company_title,omitempty" xml:"company_title,omitempty"`
	// 个人email
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 个人手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 用户名
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 专票信息
	ValueAddedInfo *ValueAddedInfo `json:"value_added_info,omitempty" xml:"value_added_info,omitempty"`
	// 发票类型（1:普通发票；2：增值税专用发票）
	InvoiceType int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 发票属性 （0：公司；1：个人）
	InvoiceAttr int64 `json:"invoice_attr,omitempty" xml:"invoice_attr,omitempty"`
	// 淘宝用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 发票id
	InvoiceId int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
}

var poolCommonInvoiceInfo = sync.Pool{
	New: func() any {
		return new(CommonInvoiceInfo)
	},
}

// GetCommonInvoiceInfo() 从对象池中获取CommonInvoiceInfo
func GetCommonInvoiceInfo() *CommonInvoiceInfo {
	return poolCommonInvoiceInfo.Get().(*CommonInvoiceInfo)
}

// ReleaseCommonInvoiceInfo 释放CommonInvoiceInfo
func ReleaseCommonInvoiceInfo(v *CommonInvoiceInfo) {
	v.CompanyTitle = ""
	v.Email = ""
	v.Phone = ""
	v.UserNick = ""
	v.ValueAddedInfo = nil
	v.InvoiceType = 0
	v.InvoiceAttr = 0
	v.UserId = 0
	v.InvoiceId = 0
	poolCommonInvoiceInfo.Put(v)
}
