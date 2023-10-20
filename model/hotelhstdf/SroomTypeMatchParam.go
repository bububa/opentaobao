package hotelhstdf

import (
	"sync"
)

// SroomTypeMatchParam 结构体
type SroomTypeMatchParam struct {
	// 标准房型id
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// 卖家房型id
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
}

var poolSroomTypeMatchParam = sync.Pool{
	New: func() any {
		return new(SroomTypeMatchParam)
	},
}

// GetSroomTypeMatchParam() 从对象池中获取SroomTypeMatchParam
func GetSroomTypeMatchParam() *SroomTypeMatchParam {
	return poolSroomTypeMatchParam.Get().(*SroomTypeMatchParam)
}

// ReleaseSroomTypeMatchParam 释放SroomTypeMatchParam
func ReleaseSroomTypeMatchParam(v *SroomTypeMatchParam) {
	v.Srid = 0
	v.Rid = 0
	poolSroomTypeMatchParam.Put(v)
}
