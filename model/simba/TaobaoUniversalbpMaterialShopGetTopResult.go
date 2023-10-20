package simba

import (
	"sync"
)

// TaobaoUniversalbpMaterialShopGetTopResult 结构体
type TaobaoUniversalbpMaterialShopGetTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	ShopInfoVO *ShopInfoVo `json:"shop_info_v_o,omitempty" xml:"shop_info_v_o,omitempty"`
}

var poolTaobaoUniversalbpMaterialShopGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpMaterialShopGetTopResult)
	},
}

// GetTaobaoUniversalbpMaterialShopGetTopResult() 从对象池中获取TaobaoUniversalbpMaterialShopGetTopResult
func GetTaobaoUniversalbpMaterialShopGetTopResult() *TaobaoUniversalbpMaterialShopGetTopResult {
	return poolTaobaoUniversalbpMaterialShopGetTopResult.Get().(*TaobaoUniversalbpMaterialShopGetTopResult)
}

// ReleaseTaobaoUniversalbpMaterialShopGetTopResult 释放TaobaoUniversalbpMaterialShopGetTopResult
func ReleaseTaobaoUniversalbpMaterialShopGetTopResult(v *TaobaoUniversalbpMaterialShopGetTopResult) {
	v.Info = nil
	v.ShopInfoVO = nil
	poolTaobaoUniversalbpMaterialShopGetTopResult.Put(v)
}
