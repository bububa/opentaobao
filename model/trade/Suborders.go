package trade

// Suborders 
/* model for simplify = false
type Suborders struct {

    // 销售单位（非标品）
    
    SaleUnit   string `json:"sale_unit,omitempty"`
    

    // 订单履约状态
    
    OrderFulfillStatus   string `json:"order_fulfill_status,omitempty"`
    

    // 销售数量
    
    SaleQuantity   int64 `json:"sale_quantity,omitempty"`
    

    // sku名称
    
    SkuName   string `json:"sku_name,omitempty"`
    

    // 销售单价
    
    SalePrice   int64 `json:"sale_price,omitempty"`
    

    // 业务订单标识
    
    BizOrderId   string `json:"biz_order_id,omitempty"`
    

    // 订单状态
    
    OrderStatus   string `json:"order_status,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 是否加工
    
    Handling   bool `json:"handling,omitempty"`
    

    // sku编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 加工方式
    
    HandlingType   string `json:"handling_type,omitempty"`
    

    // 外部关联订单标识
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 订单金额
    
    OriginFee   int64 `json:"origin_fee,omitempty"`
    

    // 取消时的作业节点:WAREHOUSE和DELIVERY_DOCK
    
    CancelNodeType   string `json:"cancel_node_type,omitempty"`
    

    // 取消结果: SUCCESS-已拦截，FAILURE-未拦截，SYSTEM_ERROR-系统异常，PARAM_ERROR-参数错误，BUSINESS_ERROR-业务异常
    
    CancelResultCode   string `json:"cancel_result_code,omitempty"`
    

    // 销售实物库存数量
    
    SaleStockQuantity   string `json:"sale_stock_quantity,omitempty"`
    

    // 实物库存单位
    
    StockUnit   string `json:"stock_unit,omitempty"`
    

    // 子商品优惠金额(可以为0)
    
    DiscountFee   int64 `json:"discount_fee,omitempty"`
    

    // 子商品实际支付金额(=商品原总-优惠金额价)
    
    PayFee   int64 `json:"pay_fee,omitempty"`
    

    // 门店编码
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 拣货数量
    
    PickSaleQuantity   int64 `json:"pick_sale_quantity,omitempty"`
    

    // 拣货金额
    
    PickStockFee   int64 `json:"pick_stock_fee,omitempty"`
    

    // 拣货实物库存数量
    
    PickSaleStockQuantity   string `json:"pick_sale_stock_quantity,omitempty"`
    

}
*/

// Suborders 
type Suborders struct {

    // 销售单位（非标品）
    SaleUnit   string `json:"sale_unit,omitempty"`

    // 订单履约状态
    OrderFulfillStatus   string `json:"order_fulfill_status,omitempty"`

    // 销售数量
    SaleQuantity   int64 `json:"sale_quantity,omitempty"`

    // sku名称
    SkuName   string `json:"sku_name,omitempty"`

    // 销售单价
    SalePrice   int64 `json:"sale_price,omitempty"`

    // 业务订单标识
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 订单状态
    OrderStatus   string `json:"order_status,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 是否加工
    Handling   bool `json:"handling,omitempty"`

    // sku编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 加工方式
    HandlingType   string `json:"handling_type,omitempty"`

    // 外部关联订单标识
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 订单金额
    OriginFee   int64 `json:"origin_fee,omitempty"`

    // 取消时的作业节点:WAREHOUSE和DELIVERY_DOCK
    CancelNodeType   string `json:"cancel_node_type,omitempty"`

    // 取消结果: SUCCESS-已拦截，FAILURE-未拦截，SYSTEM_ERROR-系统异常，PARAM_ERROR-参数错误，BUSINESS_ERROR-业务异常
    CancelResultCode   string `json:"cancel_result_code,omitempty"`

    // 销售实物库存数量
    SaleStockQuantity   string `json:"sale_stock_quantity,omitempty"`

    // 实物库存单位
    StockUnit   string `json:"stock_unit,omitempty"`

    // 子商品优惠金额(可以为0)
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 子商品实际支付金额(=商品原总-优惠金额价)
    PayFee   int64 `json:"pay_fee,omitempty"`

    // 门店编码
    ShopId   string `json:"shop_id,omitempty"`

    // 拣货数量
    PickSaleQuantity   int64 `json:"pick_sale_quantity,omitempty"`

    // 拣货金额
    PickStockFee   int64 `json:"pick_stock_fee,omitempty"`

    // 拣货实物库存数量
    PickSaleStockQuantity   string `json:"pick_sale_stock_quantity,omitempty"`

}
