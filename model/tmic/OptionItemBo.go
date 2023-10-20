package tmic

import (
	"sync"
)

// OptionItemBo 结构体
type OptionItemBo struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// tip
	Tip string `json:"tip,omitempty" xml:"tip,omitempty"`
	// randomGroupNumber
	RandomGroupNumber int64 `json:"random_group_number,omitempty" xml:"random_group_number,omitempty"`
	// supplement
	Supplement bool `json:"supplement,omitempty" xml:"supplement,omitempty"`
	// exclusion
	Exclusion bool `json:"exclusion,omitempty" xml:"exclusion,omitempty"`
	// end
	End bool `json:"end,omitempty" xml:"end,omitempty"`
}

var poolOptionItemBo = sync.Pool{
	New: func() any {
		return new(OptionItemBo)
	},
}

// GetOptionItemBo() 从对象池中获取OptionItemBo
func GetOptionItemBo() *OptionItemBo {
	return poolOptionItemBo.Get().(*OptionItemBo)
}

// ReleaseOptionItemBo 释放OptionItemBo
func ReleaseOptionItemBo(v *OptionItemBo) {
	v.Code = ""
	v.Value = ""
	v.Tip = ""
	v.RandomGroupNumber = 0
	v.Supplement = false
	v.Exclusion = false
	v.End = false
	poolOptionItemBo.Put(v)
}
