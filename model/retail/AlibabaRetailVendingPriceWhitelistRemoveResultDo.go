package retail

import (
	"sync"
)

// AlibabaRetailVendingPriceWhitelistRemoveResultDo 结构体
type AlibabaRetailVendingPriceWhitelistRemoveResultDo struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 成功标识
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolAlibabaRetailVendingPriceWhitelistRemoveResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaRetailVendingPriceWhitelistRemoveResultDo)
	},
}

// GetAlibabaRetailVendingPriceWhitelistRemoveResultDo() 从对象池中获取AlibabaRetailVendingPriceWhitelistRemoveResultDo
func GetAlibabaRetailVendingPriceWhitelistRemoveResultDo() *AlibabaRetailVendingPriceWhitelistRemoveResultDo {
	return poolAlibabaRetailVendingPriceWhitelistRemoveResultDo.Get().(*AlibabaRetailVendingPriceWhitelistRemoveResultDo)
}

// ReleaseAlibabaRetailVendingPriceWhitelistRemoveResultDo 释放AlibabaRetailVendingPriceWhitelistRemoveResultDo
func ReleaseAlibabaRetailVendingPriceWhitelistRemoveResultDo(v *AlibabaRetailVendingPriceWhitelistRemoveResultDo) {
	v.Code = ""
	v.Msg = ""
	v.Succ = false
	poolAlibabaRetailVendingPriceWhitelistRemoveResultDo.Put(v)
}
