package tmallservice

// ServiceTradeOrder 
type ServiceTradeOrder struct {

    // 服务商品的类目
    CategoryId   int64 `json:"category_id,omitempty"`

    // 服务商品的商家编码
    OuterIdSku   string `json:"outer_id_sku,omitempty"`

    // 服务商品的sku描述
    SkuDesc   string `json:"sku_desc,omitempty"`

    // 服务商品的卖家昵称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 服务商品的店铺名
    ShopName   string `json:"shop_name,omitempty"`

    // 服务商品的sku id
    SkuId   int64 `json:"sku_id,omitempty"`

    // 服务商品的id
    AuctionId   int64 `json:"auction_id,omitempty"`

    // 服务商品的标题
    AuctionTitle   string `json:"auction_title,omitempty"`

    // 付款时间
    GmtPay   string `json:"gmt_pay,omitempty"`

    // 服务商品的采购价。单位为分
    PurchasePriceUnit   int64 `json:"purchase_price_unit,omitempty"`

    // 订单id
    OrderId   int64 `json:"order_id,omitempty"`

    // 购买数量
    BuyAmount   int64 `json:"buy_amount,omitempty"`

    // 服务商品额外属性
    Attributes   string `json:"attributes,omitempty"`

}
