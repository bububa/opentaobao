package qimen

import (
	"sync"
)

// TaobaoQimenWarehouseinfoSynchronizeResponse 结构体
type TaobaoQimenWarehouseinfoSynchronizeResponse struct {
	// 仓库信息
	WarehouseInfos []WarehouseInfo `json:"warehouseInfos,omitempty" xml:"warehouseInfos>warehouse_info,omitempty"`
	// success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenWarehouseinfoSynchronizeResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWarehouseinfoSynchronizeResponse)
	},
}

// GetTaobaoQimenWarehouseinfoSynchronizeResponse() 从对象池中获取TaobaoQimenWarehouseinfoSynchronizeResponse
func GetTaobaoQimenWarehouseinfoSynchronizeResponse() *TaobaoQimenWarehouseinfoSynchronizeResponse {
	return poolTaobaoQimenWarehouseinfoSynchronizeResponse.Get().(*TaobaoQimenWarehouseinfoSynchronizeResponse)
}

// ReleaseTaobaoQimenWarehouseinfoSynchronizeResponse 释放TaobaoQimenWarehouseinfoSynchronizeResponse
func ReleaseTaobaoQimenWarehouseinfoSynchronizeResponse(v *TaobaoQimenWarehouseinfoSynchronizeResponse) {
	v.WarehouseInfos = v.WarehouseInfos[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenWarehouseinfoSynchronizeResponse.Put(v)
}
