package bus

import (
	"sync"
)

// AccountInDetail 结构体
type AccountInDetail struct {
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 对应该支付宝的支付宝账号ID，注意和支付宝账号保持一致
	AlipayAccountId string `json:"alipay_account_id,omitempty" xml:"alipay_account_id,omitempty"`
	// 单位分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolAccountInDetail = sync.Pool{
	New: func() any {
		return new(AccountInDetail)
	},
}

// GetAccountInDetail() 从对象池中获取AccountInDetail
func GetAccountInDetail() *AccountInDetail {
	return poolAccountInDetail.Get().(*AccountInDetail)
}

// ReleaseAccountInDetail 释放AccountInDetail
func ReleaseAccountInDetail(v *AccountInDetail) {
	v.AlipayAccount = ""
	v.AlipayAccountId = ""
	v.Amount = 0
	poolAccountInDetail.Put(v)
}
