package tbk

import (
	"sync"
)

// ShopUrlList 结构体
type ShopUrlList struct {
	// 物料对应错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 入参的店铺ID
	InputShopId string `json:"input_shop_id,omitempty" xml:"input_shop_id,omitempty"`
	// 物料对应错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 链接信息
	LinkInfoDto *LinkInfoDto `json:"link_info_dto,omitempty" xml:"link_info_dto,omitempty"`
}

var poolShopUrlList = sync.Pool{
	New: func() any {
		return new(ShopUrlList)
	},
}

// GetShopUrlList() 从对象池中获取ShopUrlList
func GetShopUrlList() *ShopUrlList {
	return poolShopUrlList.Get().(*ShopUrlList)
}

// ReleaseShopUrlList 释放ShopUrlList
func ReleaseShopUrlList(v *ShopUrlList) {
	v.Msg = ""
	v.InputShopId = ""
	v.Code = 0
	v.LinkInfoDto = nil
	poolShopUrlList.Put(v)
}
