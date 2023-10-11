package tmallservice

// MasterTradeOrder 结构体
type MasterTradeOrder struct {
	// 签收时间
	GmtArrival string `json:"gmt_arrival,omitempty" xml:"gmt_arrival,omitempty"`
	// 预计签收时间
	GmtExpectArrival string `json:"gmt_expect_arrival,omitempty" xml:"gmt_expect_arrival,omitempty"`
	// 商家编码
	OuterIdSku string `json:"outer_id_sku,omitempty" xml:"outer_id_sku,omitempty"`
	// sku描述
	SkuDesc string `json:"sku_desc,omitempty" xml:"sku_desc,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 店铺名
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 商品标题
	AuctionTitle string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
	// 商品额外属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 发货时间
	GmtShipped string `json:"gmt_shipped,omitempty" xml:"gmt_shipped,omitempty"`
	// 付款时间
	GmtPay string `json:"gmt_pay,omitempty" xml:"gmt_pay,omitempty"`
	// 商品编码
	OuterIdP string `json:"outer_id_p,omitempty" xml:"outer_id_p,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// spu id
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 价格。单位为分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品id
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 实物主订单号
	ParentOrderId int64 `json:"parent_order_id,omitempty" xml:"parent_order_id,omitempty"`
}
