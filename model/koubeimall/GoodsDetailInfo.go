package koubeimall

import (
	"sync"
)

// GoodsDetailInfo 结构体
type GoodsDetailInfo struct {
	// 菜品标签
	GoodsDetailTags []string `json:"goods_detail_tags,omitempty" xml:"goods_detail_tags>string,omitempty"`
	// 菜品描述
	GoodsDetailDesc string `json:"goods_detail_desc,omitempty" xml:"goods_detail_desc,omitempty"`
	// 菜品质量分
	GoodsDetailScore string `json:"goods_detail_score,omitempty" xml:"goods_detail_score,omitempty"`
	// 展示商品价格
	GoodsDetailPrice string `json:"goods_detail_price,omitempty" xml:"goods_detail_price,omitempty"`
	// 展示图片名称
	GoodsDetailName string `json:"goods_detail_name,omitempty" xml:"goods_detail_name,omitempty"`
	// 图片模型
	Picture *Picture `json:"picture,omitempty" xml:"picture,omitempty"`
}

var poolGoodsDetailInfo = sync.Pool{
	New: func() any {
		return new(GoodsDetailInfo)
	},
}

// GetGoodsDetailInfo() 从对象池中获取GoodsDetailInfo
func GetGoodsDetailInfo() *GoodsDetailInfo {
	return poolGoodsDetailInfo.Get().(*GoodsDetailInfo)
}

// ReleaseGoodsDetailInfo 释放GoodsDetailInfo
func ReleaseGoodsDetailInfo(v *GoodsDetailInfo) {
	v.GoodsDetailTags = v.GoodsDetailTags[:0]
	v.GoodsDetailDesc = ""
	v.GoodsDetailScore = ""
	v.GoodsDetailPrice = ""
	v.GoodsDetailName = ""
	v.Picture = nil
	poolGoodsDetailInfo.Put(v)
}
