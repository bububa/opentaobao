package wlb

import (
	"sync"
)

// BmsConsignOrderConfirm 结构体
type BmsConsignOrderConfirm struct {
	// 运单信息列表
	TmsOrders []TmsOrders `json:"tms_orders,omitempty" xml:"tms_orders>tms_orders,omitempty"`
	// 订单商品信息列表
	OrderItems []OrderItems `json:"order_items,omitempty" xml:"order_items>order_items,omitempty"`
	// 仓库出库时间
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 仓库作业单号，LBX号
	StoreOrderCode string `json:"store_order_code,omitempty" xml:"store_order_code,omitempty"`
	// 发货仓的仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 店铺id，主店铺时跟货主id相同
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// out_biz_id，（非导入时为订单创建时的交易号）
	ErpOrderCode string `json:"erp_order_code,omitempty" xml:"erp_order_code,omitempty"`
	// 每次发货均重新生成
	ConsignCode string `json:"consign_code,omitempty" xml:"consign_code,omitempty"`
	// BMS订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 货主id
	OwnerUserId string `json:"owner_user_id,omitempty" xml:"owner_user_id,omitempty"`
	// 邮费，以分为单位
	OrderPostFee int64 `json:"order_post_fee,omitempty" xml:"order_post_fee,omitempty"`
	// 交易订单金额，以分为单位
	OrderAmount int64 `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// 0销售平台、1手工录入、2导入发货、3ERP推送
	OrderSoruce int64 `json:"order_soruce,omitempty" xml:"order_soruce,omitempty"`
	// 操作子类型(201 一般交易出库单,502 换货出库单,503 补发出库单)
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
}

var poolBmsConsignOrderConfirm = sync.Pool{
	New: func() any {
		return new(BmsConsignOrderConfirm)
	},
}

// GetBmsConsignOrderConfirm() 从对象池中获取BmsConsignOrderConfirm
func GetBmsConsignOrderConfirm() *BmsConsignOrderConfirm {
	return poolBmsConsignOrderConfirm.Get().(*BmsConsignOrderConfirm)
}

// ReleaseBmsConsignOrderConfirm 释放BmsConsignOrderConfirm
func ReleaseBmsConsignOrderConfirm(v *BmsConsignOrderConfirm) {
	v.TmsOrders = v.TmsOrders[:0]
	v.OrderItems = v.OrderItems[:0]
	v.ConsignTime = ""
	v.StoreOrderCode = ""
	v.StoreCode = ""
	v.ShopId = ""
	v.ErpOrderCode = ""
	v.ConsignCode = ""
	v.OrderCode = ""
	v.OwnerUserId = ""
	v.OrderPostFee = 0
	v.OrderAmount = 0
	v.OrderSoruce = 0
	v.OrderType = 0
	poolBmsConsignOrderConfirm.Put(v)
}
