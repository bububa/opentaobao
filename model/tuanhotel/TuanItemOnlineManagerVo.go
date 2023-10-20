package tuanhotel

import (
	"sync"
)

// TuanItemOnlineManagerVo 结构体
type TuanItemOnlineManagerVo struct {
	// sku
	TuanItemOnlineSkus []TuanItemOnlineSkuManagerVo `json:"tuan_item_online_skus,omitempty" xml:"tuan_item_online_skus>tuan_item_online_sku_manager_vo,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 下线
	OnlineStatusDesc string `json:"online_status_desc,omitempty" xml:"online_status_desc,omitempty"`
	// itemid
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 审核状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
}

var poolTuanItemOnlineManagerVo = sync.Pool{
	New: func() any {
		return new(TuanItemOnlineManagerVo)
	},
}

// GetTuanItemOnlineManagerVo() 从对象池中获取TuanItemOnlineManagerVo
func GetTuanItemOnlineManagerVo() *TuanItemOnlineManagerVo {
	return poolTuanItemOnlineManagerVo.Get().(*TuanItemOnlineManagerVo)
}

// ReleaseTuanItemOnlineManagerVo 释放TuanItemOnlineManagerVo
func ReleaseTuanItemOnlineManagerVo(v *TuanItemOnlineManagerVo) {
	v.TuanItemOnlineSkus = v.TuanItemOnlineSkus[:0]
	v.GmtCreate = ""
	v.OnlineStatusDesc = ""
	v.ItemId = 0
	v.OnlineStatus = 0
	poolTuanItemOnlineManagerVo.Put(v)
}
