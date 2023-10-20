package rhino

import (
	"sync"
)

// SkuBo 结构体
type SkuBo struct {
	// 款式id
	StyleId string `json:"style_id,omitempty" xml:"style_id,omitempty"`
	// 工艺版本id
	TechnologyVersion string `json:"technology_version,omitempty" xml:"technology_version,omitempty"`
	// 颜色id
	ColorId int64 `json:"color_id,omitempty" xml:"color_id,omitempty"`
	// 尺码id
	SizeId int64 `json:"size_id,omitempty" xml:"size_id,omitempty"`
}

var poolSkuBo = sync.Pool{
	New: func() any {
		return new(SkuBo)
	},
}

// GetSkuBo() 从对象池中获取SkuBo
func GetSkuBo() *SkuBo {
	return poolSkuBo.Get().(*SkuBo)
}

// ReleaseSkuBo 释放SkuBo
func ReleaseSkuBo(v *SkuBo) {
	v.StyleId = ""
	v.TechnologyVersion = ""
	v.ColorId = 0
	v.SizeId = 0
	poolSkuBo.Put(v)
}
