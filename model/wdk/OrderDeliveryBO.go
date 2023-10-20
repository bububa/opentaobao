package wdk

import (
	"sync"
)

// OrderDeliveryBO 结构体
type OrderDeliveryBO struct {
	// 外部子单号
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 后端商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 前端商品名称
	AuctionTitle string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 库存单位购买数量
	StockQuantity string `json:"stock_quantity,omitempty" xml:"stock_quantity,omitempty"`
	// 库存单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 创单时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 订单来源
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 下单终端
	OrderTerminal string `json:"order_terminal,omitempty" xml:"order_terminal,omitempty"`
	// 渠道店
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 期望送达时间
	ExpectArriveTime string `json:"expect_arrive_time,omitempty" xml:"expect_arrive_time,omitempty"`
	// 买家昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 收货人名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 收货人电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收货人地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// poi经纬度
	Geo string `json:"geo,omitempty" xml:"geo,omitempty"`
	// 是否是主单 1-是;0-不是
	IsMain int64 `json:"is_main,omitempty" xml:"is_main,omitempty"`
	// 是否是子单 1-是;0-不是
	IsDetail int64 `json:"is_detail,omitempty" xml:"is_detail,omitempty"`
	// 前端商品id
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 商品单价
	SalePrice int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 销售单位购买数量
	SaleQuantity int64 `json:"sale_quantity,omitempty" xml:"sale_quantity,omitempty"`
	// 原价
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 支付金额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 业务类型 2-表示生鲜门店;3-表示B2C
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 子业务类型
	SubBusinessType int64 `json:"sub_business_type,omitempty" xml:"sub_business_type,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 订单来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 订单渠道 2-表示线上;3-表示线下
	OrderChannel int64 `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 配送类型 1-表示预约配送;2-表示现场购买
	DeliverType int64 `json:"deliver_type,omitempty" xml:"deliver_type,omitempty"`
	// 送达类型 1-表示及时达;2-表示定时达;3-表示极速达;4-表示无需配送
	ArriveType int64 `json:"arrive_type,omitempty" xml:"arrive_type,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 运费(分为单位)
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 打包费(分为单位)
	PackageFee int64 `json:"package_fee,omitempty" xml:"package_fee,omitempty"`
	// 平台折扣费(分为单位)
	PlatformDiscountFee int64 `json:"platform_discount_fee,omitempty" xml:"platform_discount_fee,omitempty"`
	// 商家折扣费(分为单位)
	MerchantDiscountFee int64 `json:"merchant_discount_fee,omitempty" xml:"merchant_discount_fee,omitempty"`
	// 买家id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolOrderDeliveryBO = sync.Pool{
	New: func() any {
		return new(OrderDeliveryBO)
	},
}

// GetOrderDeliveryBO() 从对象池中获取OrderDeliveryBO
func GetOrderDeliveryBO() *OrderDeliveryBO {
	return poolOrderDeliveryBO.Get().(*OrderDeliveryBO)
}

// ReleaseOrderDeliveryBO 释放OrderDeliveryBO
func ReleaseOrderDeliveryBO(v *OrderDeliveryBO) {
	v.OutSubOrderId = ""
	v.SkuCode = ""
	v.AuctionTitle = ""
	v.SaleUnit = ""
	v.StockQuantity = ""
	v.StockUnit = ""
	v.OrderCreateTime = ""
	v.SellerNick = ""
	v.OrderSource = ""
	v.OrderTerminal = ""
	v.ShopId = ""
	v.StoreId = ""
	v.PayTime = ""
	v.ExpectArriveTime = ""
	v.UserNick = ""
	v.Name = ""
	v.Phone = ""
	v.Address = ""
	v.Geo = ""
	v.IsMain = 0
	v.IsDetail = 0
	v.AuctionId = 0
	v.SalePrice = 0
	v.SaleQuantity = 0
	v.OriginFee = 0
	v.PayFee = 0
	v.DiscountFee = 0
	v.BusinessType = 0
	v.SubBusinessType = 0
	v.SellerId = 0
	v.OrderFrom = 0
	v.OrderChannel = 0
	v.DeliverType = 0
	v.ArriveType = 0
	v.OrderStatus = 0
	v.PostFee = 0
	v.PackageFee = 0
	v.PlatformDiscountFee = 0
	v.MerchantDiscountFee = 0
	v.UserId = 0
	poolOrderDeliveryBO.Put(v)
}
