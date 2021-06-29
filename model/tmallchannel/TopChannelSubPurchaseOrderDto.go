package tmallchannel

// TopChannelSubPurchaseOrderDTO 
type TopChannelSubPurchaseOrderDTO struct {
    // 子采购单单号
    SubPurchaseOrderNo   int64 `json:"sub_purchase_order_no,omitempty" xml:"sub_purchase_order_no,omitempty"`
    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    // 产品id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // sku信息
    Sku   string `json:"sku,omitempty" xml:"sku,omitempty"`
    // 产品编码
    ProductNumber   string `json:"product_number,omitempty" xml:"product_number,omitempty"`
    // 产品标题
    ProductTitle   string `json:"product_title,omitempty" xml:"product_title,omitempty"`
    // 产品价格
    ProductPrice   int64 `json:"product_price,omitempty" xml:"product_price,omitempty"`
    // 购买数量
    BuyAmount   int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
    // 支付状态
    PayStatus   int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
    // 子采购单的物流状态
    LogisticsStatus   int64 `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
    // 交易完成时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
