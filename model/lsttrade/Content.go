package lsttrade

// Content 
type Content struct {

    // 退款数量
    
    RefundCount   int64 `json:"refund_count,omitempty" xml:"refund_count,omitempty"`
    

    // 退款描述
    
    Discription   string `json:"discription,omitempty" xml:"discription,omitempty"`
    

    // 退款原因
    
    ApplyReason   string `json:"apply_reason,omitempty" xml:"apply_reason,omitempty"`
    

    // 运费，单位分
    
    Freight   int64 `json:"freight,omitempty" xml:"freight,omitempty"`
    

    // 退款状态，等待卖家同意("waitselleragree"), 退款成功("refundsuccess"), 退款关闭("refundclose"), 待买家修改("waitbuyermodify"), 等待买家退货("waitbuyersend"), 等待卖家确认收货("waitsellerreceive");
    
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    

    // 子单
    
    SubOrders   []SubOrders `json:"sub_orders,omitempty" xml:"sub_orders,omitempty"`
    

    // 退款完成时间
    
    GmtCompleted   string `json:"gmt_completed,omitempty" xml:"gmt_completed,omitempty"`
    

    // 买家的loginId
    
    BuyerLoginId   string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
    

    // 仓库类型
    
    WarehouseType   string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
    

    // 申请退款时间
    
    GmtApply   string `json:"gmt_apply,omitempty" xml:"gmt_apply,omitempty"`
    

    // 退款单ID
    
    RefundId   string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    

    // 退款金额，单位分
    
    RefundPayment   int64 `json:"refund_payment,omitempty" xml:"refund_payment,omitempty"`
    

    // 店铺名称
    
    BuyerShopName   string `json:"buyer_shop_name,omitempty" xml:"buyer_shop_name,omitempty"`
    

    // 主订单id
    
    MainOrderId   int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    

    // 备注信息
    
    BaseInfo   *BaseInfo `json:"base_info,omitempty" xml:"base_info,omitempty"`
    

    // 仓库类信息，此字段将会废弃，请看sub_orders下lst_warehouse_type、warehouse_code、warehouse_name
    
    OrderBizInfo   *OrderBizInfo `json:"order_biz_info,omitempty" xml:"order_biz_info,omitempty"`
    

    // 扩展字段，暂不使用
    
    ExtAttributes   string `json:"ext_attributes,omitempty" xml:"ext_attributes,omitempty"`
    

}
