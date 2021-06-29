package lstlogistics2

// LstOffLineOrderUploadParam 
type LstOffLineOrderUploadParam struct {

    // 小店名称
    
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    

    // 仓库名称
    
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    

    // 订单支付金额
    
    PayAmount   int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
    

    // 外部订单号,保证在商家下唯一
    
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    

    // 订单创建时间
    
    OrderCreateTime   string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
    

    // 订单支付时间
    
    OrderPayTime   string `json:"order_pay_time,omitempty" xml:"order_pay_time,omitempty"`
    

    // 零售通小店id
    
    LstShopId   int64 `json:"lst_shop_id,omitempty" xml:"lst_shop_id,omitempty"`
    

    // erp小店id
    
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    

    // 收货人联系方式
    
    ReceiverContact   *ContactParam `json:"receiver_contact,omitempty" xml:"receiver_contact,omitempty"`
    

    // 发货人联系方式
    
    SenderContact   *ContactParam `json:"sender_contact,omitempty" xml:"sender_contact,omitempty"`
    

    // 收货地址
    
    ReceiverAddress   *AddressParam `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    

    // 子订单
    
    SubOrders   []SubOrderParam `json:"sub_orders,omitempty" xml:"sub_orders,omitempty"`
    

    // 买家备注
    
    BuyerMessage   string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
    

}
