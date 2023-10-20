package wdk

import (
	"sync"
)

// ShopDo 结构体
type ShopDo struct {
	// 门店编码(所属的OU的编码)
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 门店编码
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 状态（默认、营业中、待维护、停用）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 类型（默认、盒马鲜生店、自定义店、盒马B2C、盒马外卖店、盒马物流中心店、盒马配送站、盒马便利店）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 省份名称
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 地区名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 商户编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 门店标记（正式、测试）
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 省份编码
	ProvCode string `json:"prov_code,omitempty" xml:"prov_code,omitempty"`
	// 城市编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 地区编码
	AreaCode string `json:"area_code,omitempty" xml:"area_code,omitempty"`
}

var poolShopDo = sync.Pool{
	New: func() any {
		return new(ShopDo)
	},
}

// GetShopDo() 从对象池中获取ShopDo
func GetShopDo() *ShopDo {
	return poolShopDo.Get().(*ShopDo)
}

// ReleaseShopDo 释放ShopDo
func ReleaseShopDo(v *ShopDo) {
	v.ShopName = ""
	v.OuCode = ""
	v.Status = ""
	v.Type = ""
	v.ProvName = ""
	v.CityName = ""
	v.AreaName = ""
	v.Address = ""
	v.MerchantCode = ""
	v.Flag = ""
	v.ProvCode = ""
	v.CityCode = ""
	v.AreaCode = ""
	poolShopDo.Put(v)
}
