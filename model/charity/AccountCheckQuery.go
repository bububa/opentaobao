package charity

import (
	"sync"
)

// AccountCheckQuery 结构体
type AccountCheckQuery struct {
	// 账单列表
	BillList []BillDto `json:"bill_list,omitempty" xml:"bill_list>bill_dto,omitempty"`
	// 公益组织id
	InvoiceOrgId string `json:"invoice_org_id,omitempty" xml:"invoice_org_id,omitempty"`
	// 开票申请id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 商家id
	MerchantId int64 `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
}

var poolAccountCheckQuery = sync.Pool{
	New: func() any {
		return new(AccountCheckQuery)
	},
}

// GetAccountCheckQuery() 从对象池中获取AccountCheckQuery
func GetAccountCheckQuery() *AccountCheckQuery {
	return poolAccountCheckQuery.Get().(*AccountCheckQuery)
}

// ReleaseAccountCheckQuery 释放AccountCheckQuery
func ReleaseAccountCheckQuery(v *AccountCheckQuery) {
	v.BillList = v.BillList[:0]
	v.InvoiceOrgId = ""
	v.InvoiceId = ""
	v.MerchantId = 0
	poolAccountCheckQuery.Put(v)
}
