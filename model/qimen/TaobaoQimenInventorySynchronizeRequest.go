package qimen

import (
	"sync"
)

// TaobaoQimenInventorySynchronizeRequest 结构体
type TaobaoQimenInventorySynchronizeRequest struct {
	// 货主编码,货主编码,string(50),,
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 仓库编码,仓库编码,string(50),,
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 调整单编码,调整单编码,string(50),,
	AdjustOrderCode string `json:"adjustOrderCode,omitempty" xml:"adjustOrderCode,omitempty"`
	// 订单创建时间,订单创建时间,string(50),,
	AdjustTime string `json:"adjustTime,omitempty" xml:"adjustTime,omitempty"`
	// 库存调整类型,库存调整类型,string(50),,
	AdjustType string `json:"adjustType,omitempty" xml:"adjustType,omitempty"`
	// 备注,备注,string(50),,
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 关联订单
	RelatedOrders *RelatedOrders `json:"relatedOrders,omitempty" xml:"relatedOrders,omitempty"`
	// 商品列表
	Items *Items `json:"items,omitempty" xml:"items,omitempty"`
}

var poolTaobaoQimenInventorySynchronizeRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventorySynchronizeRequest)
	},
}

// GetTaobaoQimenInventorySynchronizeRequest() 从对象池中获取TaobaoQimenInventorySynchronizeRequest
func GetTaobaoQimenInventorySynchronizeRequest() *TaobaoQimenInventorySynchronizeRequest {
	return poolTaobaoQimenInventorySynchronizeRequest.Get().(*TaobaoQimenInventorySynchronizeRequest)
}

// ReleaseTaobaoQimenInventorySynchronizeRequest 释放TaobaoQimenInventorySynchronizeRequest
func ReleaseTaobaoQimenInventorySynchronizeRequest(v *TaobaoQimenInventorySynchronizeRequest) {
	v.OwnerCode = ""
	v.WarehouseCode = ""
	v.AdjustOrderCode = ""
	v.AdjustTime = ""
	v.AdjustType = ""
	v.Remark = ""
	v.RelatedOrders = nil
	v.Items = nil
	poolTaobaoQimenInventorySynchronizeRequest.Put(v)
}
