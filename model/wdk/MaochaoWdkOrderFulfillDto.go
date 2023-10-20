package wdk

import (
	"sync"
)

// MaochaoWdkOrderFulfillDto 结构体
type MaochaoWdkOrderFulfillDto struct {
	// 商户编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 拣货数量
	PickAmountStock string `json:"pick_amount_stock,omitempty" xml:"pick_amount_stock,omitempty"`
	// 扩展属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 履约状态
	FulfillStatus string `json:"fulfill_status,omitempty" xml:"fulfill_status,omitempty"`
	// 渠道店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 返仓数量
	EnterWarehouseStockQuantity string `json:"enter_warehouse_stock_quantity,omitempty" xml:"enter_warehouse_stock_quantity,omitempty"`
	// 主站子订单ID
	TbSubOrderId int64 `json:"tb_sub_order_id,omitempty" xml:"tb_sub_order_id,omitempty"`
	// 五道口订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 五道口子订单ID
	BizSubOrderId int64 `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// 主站订单ID
	TbOrderId int64 `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
}

var poolMaochaoWdkOrderFulfillDto = sync.Pool{
	New: func() any {
		return new(MaochaoWdkOrderFulfillDto)
	},
}

// GetMaochaoWdkOrderFulfillDto() 从对象池中获取MaochaoWdkOrderFulfillDto
func GetMaochaoWdkOrderFulfillDto() *MaochaoWdkOrderFulfillDto {
	return poolMaochaoWdkOrderFulfillDto.Get().(*MaochaoWdkOrderFulfillDto)
}

// ReleaseMaochaoWdkOrderFulfillDto 释放MaochaoWdkOrderFulfillDto
func ReleaseMaochaoWdkOrderFulfillDto(v *MaochaoWdkOrderFulfillDto) {
	v.MerchantCode = ""
	v.StoreId = ""
	v.PickAmountStock = ""
	v.Attributes = ""
	v.FulfillStatus = ""
	v.ShopId = ""
	v.EnterWarehouseStockQuantity = ""
	v.TbSubOrderId = 0
	v.BizOrderId = 0
	v.BizSubOrderId = 0
	v.TbOrderId = 0
	poolMaochaoWdkOrderFulfillDto.Put(v)
}
