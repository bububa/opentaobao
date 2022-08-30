package tmallservice

// MasterTradeOrderDto 结构体
type MasterTradeOrderDto struct {
	// 卖家名称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 商品名称
	AuctionTitle string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
	// 商品属性(json格式)
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 卖家店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 实物主订单id
	ParentBizOrderId string `json:"parent_biz_order_id,omitempty" xml:"parent_biz_order_id,omitempty"`
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// sku描述
	SkuDesc string `json:"sku_desc,omitempty" xml:"sku_desc,omitempty"`
	// 商家编码
	OuterIdSku string `json:"outer_id_sku,omitempty" xml:"outer_id_sku,omitempty"`
	// 实物子订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 商品id
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 商家价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// spu_id
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// sku_id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
