package happytrip

import (
	"sync"
)

// AlibabaHappytripTaxiProviderAccountBalanceData 结构体
type AlibabaHappytripTaxiProviderAccountBalanceData struct {
	// 管理员列表
	Management []Managers `json:"management,omitempty" xml:"management>managers,omitempty"`
	// 可用额度(单位：分)（余额+信用额度）表示可以支撑使用的金额
	Balance int64 `json:"balance,omitempty" xml:"balance,omitempty"`
	// 本月已用金额
	UsageMonth int64 `json:"usage_month,omitempty" xml:"usage_month,omitempty"`
}

var poolAlibabaHappytripTaxiProviderAccountBalanceData = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiProviderAccountBalanceData)
	},
}

// GetAlibabaHappytripTaxiProviderAccountBalanceData() 从对象池中获取AlibabaHappytripTaxiProviderAccountBalanceData
func GetAlibabaHappytripTaxiProviderAccountBalanceData() *AlibabaHappytripTaxiProviderAccountBalanceData {
	return poolAlibabaHappytripTaxiProviderAccountBalanceData.Get().(*AlibabaHappytripTaxiProviderAccountBalanceData)
}

// ReleaseAlibabaHappytripTaxiProviderAccountBalanceData 释放AlibabaHappytripTaxiProviderAccountBalanceData
func ReleaseAlibabaHappytripTaxiProviderAccountBalanceData(v *AlibabaHappytripTaxiProviderAccountBalanceData) {
	v.Management = v.Management[:0]
	v.Balance = 0
	v.UsageMonth = 0
	poolAlibabaHappytripTaxiProviderAccountBalanceData.Put(v)
}
