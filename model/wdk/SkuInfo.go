package wdk

// SkuInfo 
type SkuInfo struct {
    // 销售单位拣货数量(若还未开始拣货可空)
    ActualSaleQuantity   string `json:"actual_sale_quantity,omitempty" xml:"actual_sale_quantity,omitempty"`
    // 库存单位拣货数量(若还未开始拣货可空)
    ActualStockQuantity   string `json:"actual_stock_quantity,omitempty" xml:"actual_stock_quantity,omitempty"`
    // 商品编码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 履约子单ID
    FulfillSubOrderId   string `json:"fulfill_sub_order_id,omitempty" xml:"fulfill_sub_order_id,omitempty"`
    // 是否缺货
    IsShortage   bool `json:"is_shortage,omitempty" xml:"is_shortage,omitempty"`
    // 扩展属性：exchangeGoods代表换货信息；containerType代表箱子标识（正常传0、禁止开箱传1）；containerCodes代表箱码数组（不存在为空）；
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
    // 容器
    Containers   []ContainerDTO `json:"containers,omitempty" xml:"containers>container_dto,omitempty"`
    // 商品名称
    SkuName   string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
    // 商品单个价格，单位分
    SkuPrice   string `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
    // 销售单位
    SkuSaleUnit   string `json:"sku_sale_unit,omitempty" xml:"sku_sale_unit,omitempty"`
    // 销售数量
    SkuSaleQuantity   string `json:"sku_sale_quantity,omitempty" xml:"sku_sale_quantity,omitempty"`
    // 商品总价，单位分
    TotalPrice   string `json:"total_price,omitempty" xml:"total_price,omitempty"`
    // 子订单差额退款金额
    RefundAmount   string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
    // 缺货数量
    OutOfStockQuantity   string `json:"out_of_stock_quantity,omitempty" xml:"out_of_stock_quantity,omitempty"`
    // 取消金额
    CancelAmount   string `json:"cancel_amount,omitempty" xml:"cancel_amount,omitempty"`
    // 是否是标品：true（“标品”），false（“非标品”）
    IsStandardSku   bool `json:"is_standard_sku,omitempty" xml:"is_standard_sku,omitempty"`
    // 取消数量
    CancelSaleQuantity   string `json:"cancel_sale_quantity,omitempty" xml:"cancel_sale_quantity,omitempty"`
    // 缺货金额
    OutOfStockAmount   string `json:"out_of_stock_amount,omitempty" xml:"out_of_stock_amount,omitempty"`
    // 盒马交易子单号
    SubBizOrderId   string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
    // 外部交易子单号
    SubSourceOrderId   string `json:"sub_source_order_id,omitempty" xml:"sub_source_order_id,omitempty"`
    // 折扣金额
    DiscountAmount   string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
    // 库存单位
    SkuStockUnit   string `json:"sku_stock_unit,omitempty" xml:"sku_stock_unit,omitempty"`
}
