package tuanhotel

import (
	"sync"
)

// RelatedPoiDetailVo 结构体
type RelatedPoiDetailVo struct {
	// POI名称
	PoiName string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	// POI ID
	PoiId int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

var poolRelatedPoiDetailVo = sync.Pool{
	New: func() any {
		return new(RelatedPoiDetailVo)
	},
}

// GetRelatedPoiDetailVo() 从对象池中获取RelatedPoiDetailVo
func GetRelatedPoiDetailVo() *RelatedPoiDetailVo {
	return poolRelatedPoiDetailVo.Get().(*RelatedPoiDetailVo)
}

// ReleaseRelatedPoiDetailVo 释放RelatedPoiDetailVo
func ReleaseRelatedPoiDetailVo(v *RelatedPoiDetailVo) {
	v.PoiName = ""
	v.PoiId = 0
	poolRelatedPoiDetailVo.Put(v)
}
