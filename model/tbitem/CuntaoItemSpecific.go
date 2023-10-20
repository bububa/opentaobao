package tbitem

import (
	"sync"
)

// CuntaoItemSpecific 结构体
type CuntaoItemSpecific struct {
	// 村淘商品级佣金率
	KickBackRate string `json:"kick_back_rate,omitempty" xml:"kick_back_rate,omitempty"`
}

var poolCuntaoItemSpecific = sync.Pool{
	New: func() any {
		return new(CuntaoItemSpecific)
	},
}

// GetCuntaoItemSpecific() 从对象池中获取CuntaoItemSpecific
func GetCuntaoItemSpecific() *CuntaoItemSpecific {
	return poolCuntaoItemSpecific.Get().(*CuntaoItemSpecific)
}

// ReleaseCuntaoItemSpecific 释放CuntaoItemSpecific
func ReleaseCuntaoItemSpecific(v *CuntaoItemSpecific) {
	v.KickBackRate = ""
	poolCuntaoItemSpecific.Put(v)
}
