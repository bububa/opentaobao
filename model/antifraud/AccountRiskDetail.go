package antifraud

import (
	"sync"
)

// AccountRiskDetail 结构体
type AccountRiskDetail struct {
	// 表示一个具体的用户风险详情
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolAccountRiskDetail = sync.Pool{
	New: func() any {
		return new(AccountRiskDetail)
	},
}

// GetAccountRiskDetail() 从对象池中获取AccountRiskDetail
func GetAccountRiskDetail() *AccountRiskDetail {
	return poolAccountRiskDetail.Get().(*AccountRiskDetail)
}

// ReleaseAccountRiskDetail 释放AccountRiskDetail
func ReleaseAccountRiskDetail(v *AccountRiskDetail) {
	v.Name = ""
	poolAccountRiskDetail.Put(v)
}
