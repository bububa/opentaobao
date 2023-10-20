package tmallhk

import (
	"sync"
)

// CtsCarriage 结构体
type CtsCarriage struct {
	// 托运开始时间
	Begin string `json:"begin,omitempty" xml:"begin,omitempty"`
	// 托运单号
	CarriageNo string `json:"carriage_no,omitempty" xml:"carriage_no,omitempty"`
	// 托运公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 托运结束时间
	End string `json:"end,omitempty" xml:"end,omitempty"`
}

var poolCtsCarriage = sync.Pool{
	New: func() any {
		return new(CtsCarriage)
	},
}

// GetCtsCarriage() 从对象池中获取CtsCarriage
func GetCtsCarriage() *CtsCarriage {
	return poolCtsCarriage.Get().(*CtsCarriage)
}

// ReleaseCtsCarriage 释放CtsCarriage
func ReleaseCtsCarriage(v *CtsCarriage) {
	v.Begin = ""
	v.CarriageNo = ""
	v.CompanyName = ""
	v.End = ""
	poolCtsCarriage.Put(v)
}
