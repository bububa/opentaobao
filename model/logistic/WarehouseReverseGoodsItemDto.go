package logistic

import (
	"sync"
)

// WarehouseReverseGoodsItemDto 结构体
type WarehouseReverseGoodsItemDto struct {
	// 扩展字段，JSONObject格式
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 商品名称
	AuctionName string `json:"auction_name,omitempty" xml:"auction_name,omitempty"`
	// 货主
	OwnerNick string `json:"owner_nick,omitempty" xml:"owner_nick,omitempty"`
	// 商品条码
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// 货品仓储系统编码
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 货品行ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 1=淘系子订单，2=赠品,3=未知
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 计划数量
	PlanQty int64 `json:"plan_qty,omitempty" xml:"plan_qty,omitempty"`
	// 实发数量
	ActualQty int64 `json:"actual_qty,omitempty" xml:"actual_qty,omitempty"`
	// 单价（单位：分）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 实付金额（单位：分）
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// 子订单ID
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 计划状态（1=正品；2=残品；3=部分正品）
	PlanStatus int64 `json:"plan_status,omitempty" xml:"plan_status,omitempty"`
	// 状态（1=正品；2=残品；3=部分正品；4=未确认）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolWarehouseReverseGoodsItemDto = sync.Pool{
	New: func() any {
		return new(WarehouseReverseGoodsItemDto)
	},
}

// GetWarehouseReverseGoodsItemDto() 从对象池中获取WarehouseReverseGoodsItemDto
func GetWarehouseReverseGoodsItemDto() *WarehouseReverseGoodsItemDto {
	return poolWarehouseReverseGoodsItemDto.Get().(*WarehouseReverseGoodsItemDto)
}

// ReleaseWarehouseReverseGoodsItemDto 释放WarehouseReverseGoodsItemDto
func ReleaseWarehouseReverseGoodsItemDto(v *WarehouseReverseGoodsItemDto) {
	v.Extra = ""
	v.AuctionName = ""
	v.OwnerNick = ""
	v.QrCode = ""
	v.ItemId = ""
	v.ItemCode = ""
	v.ItemName = ""
	v.Id = ""
	v.Type = 0
	v.PlanQty = 0
	v.ActualQty = 0
	v.Price = 0
	v.ActualFee = 0
	v.Oid = 0
	v.PlanStatus = 0
	v.Status = 0
	poolWarehouseReverseGoodsItemDto.Put(v)
}
