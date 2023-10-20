package logistic

import (
	"sync"
)

// WmsMaterialRequest 结构体
type WmsMaterialRequest struct {
	// 包裹明细
	Packages []WmsMaterialPackageDto `json:"packages,omitempty" xml:"packages>wms_material_package_dto,omitempty"`
	// 订单类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 仓储系统出库单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 货主编码
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 出库单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 1=包材; 2=耗材
	MaterialType int64 `json:"material_type,omitempty" xml:"material_type,omitempty"`
}

var poolWmsMaterialRequest = sync.Pool{
	New: func() any {
		return new(WmsMaterialRequest)
	},
}

// GetWmsMaterialRequest() 从对象池中获取WmsMaterialRequest
func GetWmsMaterialRequest() *WmsMaterialRequest {
	return poolWmsMaterialRequest.Get().(*WmsMaterialRequest)
}

// ReleaseWmsMaterialRequest 释放WmsMaterialRequest
func ReleaseWmsMaterialRequest(v *WmsMaterialRequest) {
	v.Packages = v.Packages[:0]
	v.OrderType = ""
	v.OrderId = ""
	v.OwnerCode = ""
	v.OrderCode = ""
	v.WarehouseCode = ""
	v.MaterialType = 0
	poolWmsMaterialRequest.Put(v)
}
