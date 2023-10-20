package store

import (
	"sync"
)

// TaobaoPlaceStorerelatesubGetT 结构体
type TaobaoPlaceStorerelatesubGetT struct {
	// 门店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门店分名称
	SubName string `json:"sub_name,omitempty" xml:"sub_name,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 省
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 门店Id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 省
	ProvCode int64 `json:"prov_code,omitempty" xml:"prov_code,omitempty"`
	// 市
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 区
	DistrictCode int64 `json:"district_code,omitempty" xml:"district_code,omitempty"`
}

var poolTaobaoPlaceStorerelatesubGetT = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStorerelatesubGetT)
	},
}

// GetTaobaoPlaceStorerelatesubGetT() 从对象池中获取TaobaoPlaceStorerelatesubGetT
func GetTaobaoPlaceStorerelatesubGetT() *TaobaoPlaceStorerelatesubGetT {
	return poolTaobaoPlaceStorerelatesubGetT.Get().(*TaobaoPlaceStorerelatesubGetT)
}

// ReleaseTaobaoPlaceStorerelatesubGetT 释放TaobaoPlaceStorerelatesubGetT
func ReleaseTaobaoPlaceStorerelatesubGetT(v *TaobaoPlaceStorerelatesubGetT) {
	v.Name = ""
	v.SubName = ""
	v.Address = ""
	v.ProvName = ""
	v.CityName = ""
	v.DistrictName = ""
	v.StoreId = 0
	v.ProvCode = 0
	v.CityCode = 0
	v.DistrictCode = 0
	poolTaobaoPlaceStorerelatesubGetT.Put(v)
}
