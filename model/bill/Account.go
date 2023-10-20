package bill

import (
	"sync"
)

// Account 结构体
type Account struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 费用科目编码
	AccountCode string `json:"account_code,omitempty" xml:"account_code,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 费用科目名称
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 费用科目编号
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 状态:0-不可用 1-可用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 费用科目类型:1-虚拟账户 2-交易 3-交易佣金 4-服务费 5-天猫积分 6-其他
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 是否订单相关:0-订单无关 1-订单相关
	RelatedOrder int64 `json:"related_order,omitempty" xml:"related_order,omitempty"`
}

var poolAccount = sync.Pool{
	New: func() any {
		return new(Account)
	},
}

// GetAccount() 从对象池中获取Account
func GetAccount() *Account {
	return poolAccount.Get().(*Account)
}

// ReleaseAccount 释放Account
func ReleaseAccount(v *Account) {
	v.GmtCreate = ""
	v.AccountCode = ""
	v.GmtModified = ""
	v.AccountName = ""
	v.AccountId = 0
	v.Status = 0
	v.AccountType = 0
	v.RelatedOrder = 0
	poolAccount.Put(v)
}
