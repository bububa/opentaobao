package travel

import (
	"sync"
)

// CatPropInfo 结构体
type CatPropInfo struct {
	// 属性PID，调用taobao.itemprops.get取得
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 属性VID，调用taobao.itempropvalues.get取得
	Vid string `json:"vid,omitempty" xml:"vid,omitempty"`
}

var poolCatPropInfo = sync.Pool{
	New: func() any {
		return new(CatPropInfo)
	},
}

// GetCatPropInfo() 从对象池中获取CatPropInfo
func GetCatPropInfo() *CatPropInfo {
	return poolCatPropInfo.Get().(*CatPropInfo)
}

// ReleaseCatPropInfo 释放CatPropInfo
func ReleaseCatPropInfo(v *CatPropInfo) {
	v.Pid = ""
	v.Vid = ""
	poolCatPropInfo.Put(v)
}
