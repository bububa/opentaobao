package tbk

import (
	"sync"
)

// UserRuleInfo 结构体
type UserRuleInfo struct {
	// 用户对应的商品详细信息
	ItemList []TaobaoTbkCartCouponExpireUserQueryMapData `json:"item_list,omitempty" xml:"item_list>taobao_tbk_cart_coupon_expire_user_query_map_data,omitempty"`
	// 用户在TOP上的openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 每条记录离线任务生成，代表当时离线任务的时间戳
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolUserRuleInfo = sync.Pool{
	New: func() any {
		return new(UserRuleInfo)
	},
}

// GetUserRuleInfo() 从对象池中获取UserRuleInfo
func GetUserRuleInfo() *UserRuleInfo {
	return poolUserRuleInfo.Get().(*UserRuleInfo)
}

// ReleaseUserRuleInfo 释放UserRuleInfo
func ReleaseUserRuleInfo(v *UserRuleInfo) {
	v.ItemList = v.ItemList[:0]
	v.OpenId = ""
	v.Version = 0
	poolUserRuleInfo.Put(v)
}
