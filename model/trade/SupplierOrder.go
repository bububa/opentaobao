package trade

// SupplierOrder 
/* model for simplify = false
type SupplierOrder struct {

    // 退款金额
    
    RefundFee   int64 `json:"refund_fee,omitempty"`
    

    // 外部门店ID
    
    OuterStoreId   string `json:"outer_store_id,omitempty"`
    

    // 买家账号
    
    BuyerNick   string `json:"buyer_nick,omitempty"`
    

    // 商品名称
    
    ItemTitle   string `json:"item_title,omitempty"`
    

    // 城市
    
    City   string `json:"city,omitempty"`
    

    // 商品总价，单位为分
    
    ItemTotalPrice   int64 `json:"item_total_price,omitempty"`
    

    // 退款状态，有两种状态，已退款和未退款
    
    RefundStatus   string `json:"refund_status,omitempty"`
    

    // 交易完成时间
    
    TradeEndTime   string `json:"trade_end_time,omitempty"`
    

    // 买家ID
    
    BuyerId   string `json:"buyer_id,omitempty"`
    

    // 交易创建时间
    
    TradeCreateTime   string `json:"trade_create_time,omitempty"`
    

    // 购买数量
    
    BuyAmount   int64 `json:"buy_amount,omitempty"`
    

    // 子订单ID
    
    SubOrderId   string `json:"sub_order_id,omitempty"`
    

    // 外部商品ID,对RT来说就是货号
    
    OuterItemId   string `json:"outer_item_id,omitempty"`
    

    // 交易状态
    
    TradeStatus   string `json:"trade_status,omitempty"`
    

    // 供货商身份标识
    
    Supplier   string `json:"supplier,omitempty"`
    

    // 退款完成时间
    
    RefundEndTime   string `json:"refund_end_time,omitempty"`
    

    // 商品价格，单位为分
    
    ItemPrice   int64 `json:"item_price,omitempty"`
    

    // 驿站名称
    
    StationName   string `json:"station_name,omitempty"`
    

    // 门店名称
    
    StoreName   string `json:"store_name,omitempty"`
    

    // 驿站ID
    
    StationId   int64 `json:"station_id,omitempty"`
    

    // 主订单ID
    
    MainOrderId   string `json:"main_order_id,omitempty"`
    

    // 营销活动开始时间
    
    ActivityStartTime   string `json:"activity_start_time,omitempty"`
    

    // 营销活动扩展属性，可能包含到货时间
    
    ActivityAttributes   string `json:"activity_attributes,omitempty"`
    

    // 营销活动类型
    
    ActivityType   string `json:"activity_type,omitempty"`
    

    // 营销活动名称
    
    ActivityName   string `json:"activity_name,omitempty"`
    

    // 营销活动id
    
    ActivityId   int64 `json:"activity_id,omitempty"`
    

    // 营销活动结束时间
    
    ActivityEndTime   string `json:"activity_end_time,omitempty"`
    

    // 实付金额，单位为分
    
    ActualTotalFee   int64 `json:"actual_total_fee,omitempty"`
    

    // 更新时间
    
    ModifiedTime   string `json:"modified_time,omitempty"`
    

}
*/

// SupplierOrder 
type SupplierOrder struct {

    // 退款金额
    RefundFee   int64 `json:"refund_fee,omitempty"`

    // 外部门店ID
    OuterStoreId   string `json:"outer_store_id,omitempty"`

    // 买家账号
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 商品名称
    ItemTitle   string `json:"item_title,omitempty"`

    // 城市
    City   string `json:"city,omitempty"`

    // 商品总价，单位为分
    ItemTotalPrice   int64 `json:"item_total_price,omitempty"`

    // 退款状态，有两种状态，已退款和未退款
    RefundStatus   string `json:"refund_status,omitempty"`

    // 交易完成时间
    TradeEndTime   string `json:"trade_end_time,omitempty"`

    // 买家ID
    BuyerId   string `json:"buyer_id,omitempty"`

    // 交易创建时间
    TradeCreateTime   string `json:"trade_create_time,omitempty"`

    // 购买数量
    BuyAmount   int64 `json:"buy_amount,omitempty"`

    // 子订单ID
    SubOrderId   string `json:"sub_order_id,omitempty"`

    // 外部商品ID,对RT来说就是货号
    OuterItemId   string `json:"outer_item_id,omitempty"`

    // 交易状态
    TradeStatus   string `json:"trade_status,omitempty"`

    // 供货商身份标识
    Supplier   string `json:"supplier,omitempty"`

    // 退款完成时间
    RefundEndTime   string `json:"refund_end_time,omitempty"`

    // 商品价格，单位为分
    ItemPrice   int64 `json:"item_price,omitempty"`

    // 驿站名称
    StationName   string `json:"station_name,omitempty"`

    // 门店名称
    StoreName   string `json:"store_name,omitempty"`

    // 驿站ID
    StationId   int64 `json:"station_id,omitempty"`

    // 主订单ID
    MainOrderId   string `json:"main_order_id,omitempty"`

    // 营销活动开始时间
    ActivityStartTime   string `json:"activity_start_time,omitempty"`

    // 营销活动扩展属性，可能包含到货时间
    ActivityAttributes   string `json:"activity_attributes,omitempty"`

    // 营销活动类型
    ActivityType   string `json:"activity_type,omitempty"`

    // 营销活动名称
    ActivityName   string `json:"activity_name,omitempty"`

    // 营销活动id
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 营销活动结束时间
    ActivityEndTime   string `json:"activity_end_time,omitempty"`

    // 实付金额，单位为分
    ActualTotalFee   int64 `json:"actual_total_fee,omitempty"`

    // 更新时间
    ModifiedTime   string `json:"modified_time,omitempty"`

}
