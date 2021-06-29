package tmallservice

// ServiceTradeOrderDTO 
type ServiceTradeOrderDTO struct {
    // 卖家名称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 商品id
    AuctionId   int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
    // 服务子订单id
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 商品名称
    AuctionTitle   string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
    // 商品价格
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 服务主订单id
    ParentBizOrderId   int64 `json:"parent_biz_order_id,omitempty" xml:"parent_biz_order_id,omitempty"`
    // 服务订单属性(json格式)
    Attribute   string `json:"attribute,omitempty" xml:"attribute,omitempty"`
    // 卖家店铺名称
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    // sku描述
    SkuDesc   string `json:"sku_desc,omitempty" xml:"sku_desc,omitempty"`
    // 类目id
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    // sku_id
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 商家编码
    OuterIdSku   string `json:"outer_id_sku,omitempty" xml:"outer_id_sku,omitempty"`
    // 服务采购价
    PurchasePriceUnit   int64 `json:"purchase_price_unit,omitempty" xml:"purchase_price_unit,omitempty"`
}
