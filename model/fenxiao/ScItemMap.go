package fenxiao

import (
	"sync"
)

// ScItemMap 结构体
type ScItemMap struct {
	// 后端商品所有者名称
	RelUserNick string `json:"rel_user_nick,omitempty" xml:"rel_user_nick,omitempty"`
	// Ic商家nick(分销商nick)
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 后端商品所有者商家编码
	RelOuterCode string `json:"rel_outer_code,omitempty" xml:"rel_outer_code,omitempty"`
	// 后端商品ID
	RelItemId int64 `json:"rel_item_id,omitempty" xml:"rel_item_id,omitempty"`
	// map_type=1时，item_id为IC商品id&lt;br/&gt;&lt;br/&gt;map_type=7时，item_id为分销商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 后端商品所有者id
	RelUserId int64 `json:"rel_user_id,omitempty" xml:"rel_user_id,omitempty"`
	// 1:前台和后台关系&lt;br/&gt;7:分销和后台关系
	MapType int64 `json:"map_type,omitempty" xml:"map_type,omitempty"`
	// Ic商家id(分销商id)
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 当宝贝下没SKU时该字段为空
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolScItemMap = sync.Pool{
	New: func() any {
		return new(ScItemMap)
	},
}

// GetScItemMap() 从对象池中获取ScItemMap
func GetScItemMap() *ScItemMap {
	return poolScItemMap.Get().(*ScItemMap)
}

// ReleaseScItemMap 释放ScItemMap
func ReleaseScItemMap(v *ScItemMap) {
	v.RelUserNick = ""
	v.UserNick = ""
	v.RelOuterCode = ""
	v.RelItemId = 0
	v.ItemId = 0
	v.RelUserId = 0
	v.MapType = 0
	v.UserId = 0
	v.SkuId = 0
	poolScItemMap.Put(v)
}
