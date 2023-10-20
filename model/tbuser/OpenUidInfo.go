package tbuser

import (
	"sync"
)

// OpenUidInfo 结构体
type OpenUidInfo struct {
	// 买家openuid
	BuyerOpenUid string `json:"buyer_open_uid,omitempty" xml:"buyer_open_uid,omitempty"`
	// 买家nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolOpenUidInfo = sync.Pool{
	New: func() any {
		return new(OpenUidInfo)
	},
}

// GetOpenUidInfo() 从对象池中获取OpenUidInfo
func GetOpenUidInfo() *OpenUidInfo {
	return poolOpenUidInfo.Get().(*OpenUidInfo)
}

// ReleaseOpenUidInfo 释放OpenUidInfo
func ReleaseOpenUidInfo(v *OpenUidInfo) {
	v.BuyerOpenUid = ""
	v.BuyerNick = ""
	v.Tid = 0
	poolOpenUidInfo.Put(v)
}
