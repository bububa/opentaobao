package qimen

import (
	"sync"
)

// TaobaoQimenWarehouseinfoQueryResponse 结构体
type TaobaoQimenWarehouseinfoQueryResponse struct {
	// 奇门仓储字段
	WarehouseInfos []WarehouseInfo `json:"warehouseInfos,omitempty" xml:"warehouseInfos>warehouse_info,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	OwnerName string `json:"ownerName,omitempty" xml:"ownerName,omitempty"`
}

var poolTaobaoQimenWarehouseinfoQueryResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWarehouseinfoQueryResponse)
	},
}

// GetTaobaoQimenWarehouseinfoQueryResponse() 从对象池中获取TaobaoQimenWarehouseinfoQueryResponse
func GetTaobaoQimenWarehouseinfoQueryResponse() *TaobaoQimenWarehouseinfoQueryResponse {
	return poolTaobaoQimenWarehouseinfoQueryResponse.Get().(*TaobaoQimenWarehouseinfoQueryResponse)
}

// ReleaseTaobaoQimenWarehouseinfoQueryResponse 释放TaobaoQimenWarehouseinfoQueryResponse
func ReleaseTaobaoQimenWarehouseinfoQueryResponse(v *TaobaoQimenWarehouseinfoQueryResponse) {
	v.WarehouseInfos = v.WarehouseInfos[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.OwnerCode = ""
	v.OwnerName = ""
	poolTaobaoQimenWarehouseinfoQueryResponse.Put(v)
}
