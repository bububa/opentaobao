package wdk

// OrderSyncRefundDto 
/* model for simplify = false
type OrderSyncRefundDto struct {

    // 库存单位购买数量
    
    BuyAmountStock   string `json:"buy_amount_stock,omitempty"`
    

    // 商品skucode
    
    ItemCode   string `json:"item_code,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 非标品购买数量
    
    NsQuantity   string `json:"ns_quantity,omitempty"`
    

    // 操作人id
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 操作人名称
    
    OperatorName   string `json:"operator_name,omitempty"`
    

    // 盒马子订单号
    
    OriginOrderId   int64 `json:"origin_order_id,omitempty"`
    

    // 盒马主订单号
    
    OriginParentId   int64 `json:"origin_parent_id,omitempty"`
    

    // 库存单位拣货数量
    
    PickAmountStock   string `json:"pick_amount_stock,omitempty"`
    

    // 原购买数量
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 退款金额
    
    RefundAmount   int64 `json:"refund_amount,omitempty"`
    

    // 退款渠道列表
    
    RefundChannelList  struct {
        OrderSyncRefundChannel  []OrderSyncRefundChannel `json:"order_sync_refund_channel,omitempty"`
    } `json:"refund_channel_list,omitempty"`
    

    // 退款单id
    
    RefundOrderId   int64 `json:"refund_order_id,omitempty"`
    

    // 退货数量
    
    RefundQuantity   string `json:"refund_quantity,omitempty"`
    

    // 退款时间
    
    RefundTime   string `json:"refund_time,omitempty"`
    

    // 退款类型
    
    RefundType   int64 `json:"refund_type,omitempty"`
    

    // 非标品单位
    
    SellUnit   string `json:"sell_unit,omitempty"`
    

    // 店铺码
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 同步状态
    
    SyncStatus   string `json:"sync_status,omitempty"`
    

    // 淘系子订单号
    
    TbBizOrderId   int64 `json:"tb_biz_order_id,omitempty"`
    

    // 跑批时间
    
    BatchTime   string `json:"batch_time,omitempty"`
    

    // 实际取货数量
    
    ActualPickQuantity   string `json:"actual_pick_quantity,omitempty"`
    

    // 实际退货数量
    
    ActualRefundQuantity   string `json:"actual_refund_quantity,omitempty"`
    

    // 履约状态
    
    PromiseStatus   string `json:"promise_status,omitempty"`
    

    // 退款运费，单位为分
    
    RefundPostFee   int64 `json:"refund_post_fee,omitempty"`
    

    // stockUnit
    
    StockUnit   string `json:"stock_unit,omitempty"`
    

    // saleUnit
    
    SaleUnit   string `json:"sale_unit,omitempty"`
    

    // 淘系主订单号
    
    TbBizParentId   int64 `json:"tb_biz_parent_id,omitempty"`
    

    // 附加属性
    
    Attributes   string `json:"attributes,omitempty"`
    

    // 交易属性
    
    TradeAttributes   string `json:"trade_attributes,omitempty"`
    

    // 商品skucode
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // open_uid
    
    OpenUid   string `json:"open_uid,omitempty"`
    

    // 2:app 3:pos
    
    OrderClient   int64 `json:"order_client,omitempty"`
    

    // 渠道来源 3：饿了么 4：盒马
    
    OrderFrom   int64 `json:"order_from,omitempty"`
    

    // 外部子订单号
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 外部主订单号
    
    OutMianOrderId   string `json:"out_mian_order_id,omitempty"`
    

    // 渠道店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 外部skuCode
    
    OutSkuCode   string `json:"out_sku_code,omitempty"`
    

    // 纠纷类型 1：售中  2：售后
    
    DisputeType   int64 `json:"dispute_type,omitempty"`
    

    // 外部退款id
    
    OutRefundId   string `json:"out_refund_id,omitempty"`
    

    // 外部门店id
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 扩展属性map
    
    RefundAttributes   string `json:"refund_attributes,omitempty"`
    

    // 供应商code
    
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`
    

}
*/

// OrderSyncRefundDto 
type OrderSyncRefundDto struct {

    // 库存单位购买数量
    BuyAmountStock   string `json:"buy_amount_stock,omitempty"`

    // 商品skucode
    ItemCode   string `json:"item_code,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 非标品购买数量
    NsQuantity   string `json:"ns_quantity,omitempty"`

    // 操作人id
    OperatorId   string `json:"operator_id,omitempty"`

    // 操作人名称
    OperatorName   string `json:"operator_name,omitempty"`

    // 盒马子订单号
    OriginOrderId   int64 `json:"origin_order_id,omitempty"`

    // 盒马主订单号
    OriginParentId   int64 `json:"origin_parent_id,omitempty"`

    // 库存单位拣货数量
    PickAmountStock   string `json:"pick_amount_stock,omitempty"`

    // 原购买数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 退款金额
    RefundAmount   int64 `json:"refund_amount,omitempty"`

    // 退款渠道列表
    RefundChannelList   []OrderSyncRefundChannel `json:"refund_channel_list,omitempty"`

    // 退款单id
    RefundOrderId   int64 `json:"refund_order_id,omitempty"`

    // 退货数量
    RefundQuantity   string `json:"refund_quantity,omitempty"`

    // 退款时间
    RefundTime   string `json:"refund_time,omitempty"`

    // 退款类型
    RefundType   int64 `json:"refund_type,omitempty"`

    // 非标品单位
    SellUnit   string `json:"sell_unit,omitempty"`

    // 店铺码
    StoreId   string `json:"store_id,omitempty"`

    // 同步状态
    SyncStatus   string `json:"sync_status,omitempty"`

    // 淘系子订单号
    TbBizOrderId   int64 `json:"tb_biz_order_id,omitempty"`

    // 跑批时间
    BatchTime   string `json:"batch_time,omitempty"`

    // 实际取货数量
    ActualPickQuantity   string `json:"actual_pick_quantity,omitempty"`

    // 实际退货数量
    ActualRefundQuantity   string `json:"actual_refund_quantity,omitempty"`

    // 履约状态
    PromiseStatus   string `json:"promise_status,omitempty"`

    // 退款运费，单位为分
    RefundPostFee   int64 `json:"refund_post_fee,omitempty"`

    // stockUnit
    StockUnit   string `json:"stock_unit,omitempty"`

    // saleUnit
    SaleUnit   string `json:"sale_unit,omitempty"`

    // 淘系主订单号
    TbBizParentId   int64 `json:"tb_biz_parent_id,omitempty"`

    // 附加属性
    Attributes   string `json:"attributes,omitempty"`

    // 交易属性
    TradeAttributes   string `json:"trade_attributes,omitempty"`

    // 商品skucode
    SkuCode   string `json:"sku_code,omitempty"`

    // open_uid
    OpenUid   string `json:"open_uid,omitempty"`

    // 2:app 3:pos
    OrderClient   int64 `json:"order_client,omitempty"`

    // 渠道来源 3：饿了么 4：盒马
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 外部子订单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 外部主订单号
    OutMianOrderId   string `json:"out_mian_order_id,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

    // 外部skuCode
    OutSkuCode   string `json:"out_sku_code,omitempty"`

    // 纠纷类型 1：售中  2：售后
    DisputeType   int64 `json:"dispute_type,omitempty"`

    // 外部退款id
    OutRefundId   string `json:"out_refund_id,omitempty"`

    // 外部门店id
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 扩展属性map
    RefundAttributes   string `json:"refund_attributes,omitempty"`

    // 供应商code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`

}
