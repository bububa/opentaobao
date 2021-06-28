package wdk

// MaochaoWdkOrderFulfillDto 
type MaochaoWdkOrderFulfillDto struct {

    // 主站子订单ID
    
    TbSubOrderId   int64 `json:"tb_sub_order_id,omitempty" xml:"tb_sub_order_id,omitempty"`
    

    // 五道口订单ID
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 五道口子订单ID
    
    BizSubOrderId   int64 `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
    

    // 商户编码
    
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    

    // 经营店ID
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 拣货数量
    
    PickAmountStock   string `json:"pick_amount_stock,omitempty" xml:"pick_amount_stock,omitempty"`
    

    // 扩展属性
    
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
    

    // 履约状态
    
    FulfillStatus   string `json:"fulfill_status,omitempty" xml:"fulfill_status,omitempty"`
    

    // 主站订单ID
    
    TbOrderId   int64 `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
    

    // 渠道店ID
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 返仓数量
    
    EnterWarehouseStockQuantity   string `json:"enter_warehouse_stock_quantity,omitempty" xml:"enter_warehouse_stock_quantity,omitempty"`
    

}
