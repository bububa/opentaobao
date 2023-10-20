package alihealthmedical

import (
	"sync"
)

// ItemPublishRequest 结构体
type ItemPublishRequest struct {
	// 医生id
	DoctorUuid string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
	// 互联网医院id
	NetHospitalId string `json:"net_hospital_id,omitempty" xml:"net_hospital_id,omitempty"`
	// 商品类目code
	CategoryCode string `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// 价格，单位：分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 每日库存上限
	StockLimit int64 `json:"stock_limit,omitempty" xml:"stock_limit,omitempty"`
}

var poolItemPublishRequest = sync.Pool{
	New: func() any {
		return new(ItemPublishRequest)
	},
}

// GetItemPublishRequest() 从对象池中获取ItemPublishRequest
func GetItemPublishRequest() *ItemPublishRequest {
	return poolItemPublishRequest.Get().(*ItemPublishRequest)
}

// ReleaseItemPublishRequest 释放ItemPublishRequest
func ReleaseItemPublishRequest(v *ItemPublishRequest) {
	v.DoctorUuid = ""
	v.NetHospitalId = ""
	v.CategoryCode = ""
	v.Price = ""
	v.ItemId = 0
	v.StockLimit = 0
	poolItemPublishRequest.Put(v)
}
