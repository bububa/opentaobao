package wdk

import (
	"sync"
)

// IsvSupplySyncParam 结构体
type IsvSupplySyncParam struct {
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 仓类型，warehouse是大仓、shop是店仓
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 仓名称
	Warehouse string `json:"warehouse,omitempty" xml:"warehouse,omitempty"`
	// 库存变动类型，arrive是到货、deliver是出货、lose是丢货
	ModifyType string `json:"modify_type,omitempty" xml:"modify_type,omitempty"`
	// 库存变动时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 订单id，用来唯一标识库存流转记录，避免重复同步
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 库存变动数量，到货是正值、出货是负值、丢货是负值
	ModifyQuantity int64 `json:"modify_quantity,omitempty" xml:"modify_quantity,omitempty"`
	// 派样活动id
	SampleActivityId int64 `json:"sample_activity_id,omitempty" xml:"sample_activity_id,omitempty"`
}

var poolIsvSupplySyncParam = sync.Pool{
	New: func() any {
		return new(IsvSupplySyncParam)
	},
}

// GetIsvSupplySyncParam() 从对象池中获取IsvSupplySyncParam
func GetIsvSupplySyncParam() *IsvSupplySyncParam {
	return poolIsvSupplySyncParam.Get().(*IsvSupplySyncParam)
}

// ReleaseIsvSupplySyncParam 释放IsvSupplySyncParam
func ReleaseIsvSupplySyncParam(v *IsvSupplySyncParam) {
	v.Barcode = ""
	v.WarehouseType = ""
	v.WarehouseCode = ""
	v.Warehouse = ""
	v.ModifyType = ""
	v.ModifyTime = ""
	v.Operator = ""
	v.OrderId = 0
	v.ModifyQuantity = 0
	v.SampleActivityId = 0
	poolIsvSupplySyncParam.Put(v)
}
