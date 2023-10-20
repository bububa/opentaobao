package axindata

import (
	"sync"
)

// PageVo 结构体
type PageVo struct {
	// 标准酒店列表
	DataList []StdHotelVo `json:"data_list,omitempty" xml:"data_list>std_hotel_vo,omitempty"`
	// 记录总条数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolPageVo = sync.Pool{
	New: func() any {
		return new(PageVo)
	},
}

// GetPageVo() 从对象池中获取PageVo
func GetPageVo() *PageVo {
	return poolPageVo.Get().(*PageVo)
}

// ReleasePageVo 释放PageVo
func ReleasePageVo(v *PageVo) {
	v.DataList = v.DataList[:0]
	v.Count = 0
	poolPageVo.Put(v)
}
