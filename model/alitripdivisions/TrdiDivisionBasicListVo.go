package alitripdivisions

import (
	"sync"
)

// TrdiDivisionBasicListVo 结构体
type TrdiDivisionBasicListVo struct {
	// model
	List []TrdiDivisionBasicVo `json:"list,omitempty" xml:"list>trdi_division_basic_vo,omitempty"`
}

var poolTrdiDivisionBasicListVo = sync.Pool{
	New: func() any {
		return new(TrdiDivisionBasicListVo)
	},
}

// GetTrdiDivisionBasicListVo() 从对象池中获取TrdiDivisionBasicListVo
func GetTrdiDivisionBasicListVo() *TrdiDivisionBasicListVo {
	return poolTrdiDivisionBasicListVo.Get().(*TrdiDivisionBasicListVo)
}

// ReleaseTrdiDivisionBasicListVo 释放TrdiDivisionBasicListVo
func ReleaseTrdiDivisionBasicListVo(v *TrdiDivisionBasicListVo) {
	v.List = v.List[:0]
	poolTrdiDivisionBasicListVo.Put(v)
}
