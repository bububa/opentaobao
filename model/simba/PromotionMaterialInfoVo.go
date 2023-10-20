package simba

import (
	"sync"
)

// PromotionMaterialInfoVo 结构体
type PromotionMaterialInfoVo struct {
	// 图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 跳转链接
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 物料名称
	MaterialName string `json:"material_name,omitempty" xml:"material_name,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 物料id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
}

var poolPromotionMaterialInfoVo = sync.Pool{
	New: func() any {
		return new(PromotionMaterialInfoVo)
	},
}

// GetPromotionMaterialInfoVo() 从对象池中获取PromotionMaterialInfoVo
func GetPromotionMaterialInfoVo() *PromotionMaterialInfoVo {
	return poolPromotionMaterialInfoVo.Get().(*PromotionMaterialInfoVo)
}

// ReleasePromotionMaterialInfoVo 释放PromotionMaterialInfoVo
func ReleasePromotionMaterialInfoVo(v *PromotionMaterialInfoVo) {
	v.ImgUrl = ""
	v.LinkUrl = ""
	v.MaterialName = ""
	v.Title = ""
	v.MaterialId = 0
	poolPromotionMaterialInfoVo.Put(v)
}
