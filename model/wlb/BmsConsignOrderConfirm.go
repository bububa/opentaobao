package wlb

// BmsConsignOrderConfirm 
/* model for simplify = false
type BmsConsignOrderConfirm struct {

    // 运单信息列表
    
    TmsOrders  struct {
        TmsOrders  []TmsOrders `json:"tms_orders,omitempty"`
    } `json:"tms_orders,omitempty"`
    

    // 订单商品信息列表
    
    OrderItems  struct {
        OrderItems  []OrderItems `json:"order_items,omitempty"`
    } `json:"order_items,omitempty"`
    

    // 邮费，以分为单位
    
    OrderPostFee   int64 `json:"order_post_fee,omitempty"`
    

    // 交易订单金额，以分为单位
    
    OrderAmount   int64 `json:"order_amount,omitempty"`
    

    // 仓库出库时间
    
    ConsignTime   string `json:"consign_time,omitempty"`
    

    // 仓库作业单号，LBX号
    
    StoreOrderCode   string `json:"store_order_code,omitempty"`
    

    // 发货仓的仓库编码
    
    StoreCode   string `json:"store_code,omitempty"`
    

    // 0销售平台、1手工录入、2导入发货、3ERP推送
    
    OrderSoruce   int64 `json:"order_soruce,omitempty"`
    

    // 店铺id，主店铺时跟货主id相同
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 操作子类型(201 一般交易出库单,502 换货出库单,503 补发出库单)
    
    OrderType   int64 `json:"order_type,omitempty"`
    

    // out_biz_id，（非导入时为订单创建时的交易号）
    
    ErpOrderCode   string `json:"erp_order_code,omitempty"`
    

    // 每次发货均重新生成
    
    ConsignCode   string `json:"consign_code,omitempty"`
    

    // BMS订单编码
    
    OrderCode   string `json:"order_code,omitempty"`
    

    // 货主id
    
    OwnerUserId   string `json:"owner_user_id,omitempty"`
    

}
*/

// BmsConsignOrderConfirm 
type BmsConsignOrderConfirm struct {

    // 运单信息列表
    TmsOrders   []TmsOrders `json:"tms_orders,omitempty"`

    // 订单商品信息列表
    OrderItems   []OrderItems `json:"order_items,omitempty"`

    // 邮费，以分为单位
    OrderPostFee   int64 `json:"order_post_fee,omitempty"`

    // 交易订单金额，以分为单位
    OrderAmount   int64 `json:"order_amount,omitempty"`

    // 仓库出库时间
    ConsignTime   string `json:"consign_time,omitempty"`

    // 仓库作业单号，LBX号
    StoreOrderCode   string `json:"store_order_code,omitempty"`

    // 发货仓的仓库编码
    StoreCode   string `json:"store_code,omitempty"`

    // 0销售平台、1手工录入、2导入发货、3ERP推送
    OrderSoruce   int64 `json:"order_soruce,omitempty"`

    // 店铺id，主店铺时跟货主id相同
    ShopId   string `json:"shop_id,omitempty"`

    // 操作子类型(201 一般交易出库单,502 换货出库单,503 补发出库单)
    OrderType   int64 `json:"order_type,omitempty"`

    // out_biz_id，（非导入时为订单创建时的交易号）
    ErpOrderCode   string `json:"erp_order_code,omitempty"`

    // 每次发货均重新生成
    ConsignCode   string `json:"consign_code,omitempty"`

    // BMS订单编码
    OrderCode   string `json:"order_code,omitempty"`

    // 货主id
    OwnerUserId   string `json:"owner_user_id,omitempty"`

}
