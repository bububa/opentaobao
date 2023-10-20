package trade

import (
	"sync"
)

// SupplierOrder 结构体
type SupplierOrder struct {
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 买家账号
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 商品名称
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 退款状态，有两种状态，已退款和未退款
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 交易完成时间
	TradeEndTime string `json:"trade_end_time,omitempty" xml:"trade_end_time,omitempty"`
	// 买家ID
	BuyerId string `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 交易创建时间
	TradeCreateTime string `json:"trade_create_time,omitempty" xml:"trade_create_time,omitempty"`
	// 子订单ID
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 外部商品ID,对RT来说就是货号
	OuterItemId string `json:"outer_item_id,omitempty" xml:"outer_item_id,omitempty"`
	// 交易状态
	TradeStatus string `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
	// 供货商身份标识
	Supplier string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	// 退款完成时间
	RefundEndTime string `json:"refund_end_time,omitempty" xml:"refund_end_time,omitempty"`
	// 驿站名称
	StationName string `json:"station_name,omitempty" xml:"station_name,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 主订单ID
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 营销活动开始时间
	ActivityStartTime string `json:"activity_start_time,omitempty" xml:"activity_start_time,omitempty"`
	// 营销活动扩展属性，可能包含到货时间
	ActivityAttributes string `json:"activity_attributes,omitempty" xml:"activity_attributes,omitempty"`
	// 营销活动类型
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 营销活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 营销活动结束时间
	ActivityEndTime string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
	// 更新时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 商品总价，单位为分
	ItemTotalPrice int64 `json:"item_total_price,omitempty" xml:"item_total_price,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 商品价格，单位为分
	ItemPrice int64 `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 驿站ID
	StationId int64 `json:"station_id,omitempty" xml:"station_id,omitempty"`
	// 营销活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 实付金额，单位为分
	ActualTotalFee int64 `json:"actual_total_fee,omitempty" xml:"actual_total_fee,omitempty"`
}

var poolSupplierOrder = sync.Pool{
	New: func() any {
		return new(SupplierOrder)
	},
}

// GetSupplierOrder() 从对象池中获取SupplierOrder
func GetSupplierOrder() *SupplierOrder {
	return poolSupplierOrder.Get().(*SupplierOrder)
}

// ReleaseSupplierOrder 释放SupplierOrder
func ReleaseSupplierOrder(v *SupplierOrder) {
	v.OuterStoreId = ""
	v.BuyerNick = ""
	v.ItemTitle = ""
	v.City = ""
	v.RefundStatus = ""
	v.TradeEndTime = ""
	v.BuyerId = ""
	v.TradeCreateTime = ""
	v.SubOrderId = ""
	v.OuterItemId = ""
	v.TradeStatus = ""
	v.Supplier = ""
	v.RefundEndTime = ""
	v.StationName = ""
	v.StoreName = ""
	v.MainOrderId = ""
	v.ActivityStartTime = ""
	v.ActivityAttributes = ""
	v.ActivityType = ""
	v.ActivityName = ""
	v.ActivityEndTime = ""
	v.ModifiedTime = ""
	v.RefundFee = 0
	v.ItemTotalPrice = 0
	v.BuyAmount = 0
	v.ItemPrice = 0
	v.StationId = 0
	v.ActivityId = 0
	v.ActualTotalFee = 0
	poolSupplierOrder.Put(v)
}
