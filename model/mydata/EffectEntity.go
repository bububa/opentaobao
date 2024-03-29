package mydata

import (
	"sync"
)

// EffectEntity 结构体
type EffectEntity struct {
	// 词来源
	KeywordEffects []KeywordEffectEntity `json:"keyword_effects,omitempty" xml:"keyword_effects>keyword_effect_entity,omitempty"`
	// 收藏买家数
	Bookmark int64 `json:"bookmark,omitempty" xml:"bookmark,omitempty"`
	// 点击
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 对比买家数
	Compare int64 `json:"compare,omitempty" xml:"compare,omitempty"`
	// 反馈
	Fb int64 `json:"fb,omitempty" xml:"fb,omitempty"`
	// 曝光
	Impression int64 `json:"impression,omitempty" xml:"impression,omitempty"`
	// 提交订单数
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 分享买家数
	Share int64 `json:"share,omitempty" xml:"share,omitempty"`
	// 访客数
	Visitor int64 `json:"visitor,omitempty" xml:"visitor,omitempty"`
}

var poolEffectEntity = sync.Pool{
	New: func() any {
		return new(EffectEntity)
	},
}

// GetEffectEntity() 从对象池中获取EffectEntity
func GetEffectEntity() *EffectEntity {
	return poolEffectEntity.Get().(*EffectEntity)
}

// ReleaseEffectEntity 释放EffectEntity
func ReleaseEffectEntity(v *EffectEntity) {
	v.KeywordEffects = v.KeywordEffects[:0]
	v.Bookmark = 0
	v.Click = 0
	v.Compare = 0
	v.Fb = 0
	v.Impression = 0
	v.Order = 0
	v.ProductId = 0
	v.Share = 0
	v.Visitor = 0
	poolEffectEntity.Put(v)
}
