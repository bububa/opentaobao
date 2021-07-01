package wdk

// MaochaoWdkOrderDto 
type MaochaoWdkOrderDto struct {
    // 订单状态
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 库存单位购买数量
    BuyAmountStock   string `json:"buy_amount_stock,omitempty" xml:"buy_amount_stock,omitempty"`
    // 经营店ID
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 销售单位
    SaleUnit   string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
    // 库存单位
    StockUnit   string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
    // 非标品售卖单位
    SellUnit   string `json:"sell_unit,omitempty" xml:"sell_unit,omitempty"`
    // 非标品购买数量
    NsQuantity   string `json:"ns_quantity,omitempty" xml:"ns_quantity,omitempty"`
    // 购买数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 商品id
    AuctionId   string `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
    // 商户编码
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    // 五道口子订单ID
    BizSubOrderId   int64 `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
    // 主站子订单ID
    TbSubOrderId   int64 `json:"tb_sub_order_id,omitempty" xml:"tb_sub_order_id,omitempty"`
    // 主站订单ID
    TbOrderId   int64 `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
    // 五道口订单ID
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 商品价格
    AuctionPrice   int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
    // 商品名称
    AuctionTitle   string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 渠道店ID
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 子订单扩展字段, 订单商品采购价数据purchase_price
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
}
