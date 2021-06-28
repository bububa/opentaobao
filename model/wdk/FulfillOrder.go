package wdk

// FulfillOrder 
type FulfillOrder struct {

    // 履约单号
    
    FulfillOrderId   string `json:"fulfill_order_id,omitempty" xml:"fulfill_order_id,omitempty"`
    

    // 子订单信息列表
    
    SkuInfoList   []SkuInfo `json:"sku_info_list,omitempty" xml:"sku_info_list,omitempty"`
    

    // 扩展属性
    
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
    

    // 收货人姓名
    
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    

    // 收货人电话
    
    BuyerPhone   string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
    

    // 收货人地址
    
    BuyerAddress   string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
    

    // 订单总金额
    
    TotalOrderAmount   string `json:"total_order_amount,omitempty" xml:"total_order_amount,omitempty"`
    

    // 订单优惠金额
    
    DiscountAmount   string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
    

    // 订单应付金额
    
    PayOrderAmount   string `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty"`
    

    // 实付金额
    
    PaidAmount   string `json:"paid_amount,omitempty" xml:"paid_amount,omitempty"`
    

    // 主订单差额退款金额
    
    RefundAmount   string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
    

    // 运费
    
    CarriageAmount   string `json:"carriage_amount,omitempty" xml:"carriage_amount,omitempty"`
    

    // 取消金额
    
    CancelAmount   string `json:"cancel_amount,omitempty" xml:"cancel_amount,omitempty"`
    

    // 缺货金额
    
    OutOfStockAmount   string `json:"out_of_stock_amount,omitempty" xml:"out_of_stock_amount,omitempty"`
    

    // 盒马交易单号
    
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 外部交易单号
    
    SourceOrderId   string `json:"source_order_id,omitempty" xml:"source_order_id,omitempty"`
    

    // 业务场景
    
    ScenarioGroup   string `json:"scenario_group,omitempty" xml:"scenario_group,omitempty"`
    

    // 订单标识，取值举例：早波次001、早波次002、早波次003
    
    OrderTag   string `json:"order_tag,omitempty" xml:"order_tag,omitempty"`
    

    // 温层标识，取值：常温、冷藏、冷冻
    
    StorageMode   string `json:"storage_mode,omitempty" xml:"storage_mode,omitempty"`
    

    // 多供给标识，取值：多1、多2、多3
    
    NewSupply   string `json:"new_supply,omitempty" xml:"new_supply,omitempty"`
    

}
