package wdk

// InsuranceOrder 
/* model for simplify = false
type InsuranceOrder struct {

    // 签收时间
    
    SignTime   string `json:"sign_time,omitempty"`
    

    // 交易单号生成时间
    
    OrderCreateTime   string `json:"order_create_time,omitempty"`
    

    // 快递单号
    
    ExpressNo   string `json:"express_no,omitempty"`
    

    // 收货地址
    
    DeliveryAddress   string `json:"delivery_address,omitempty"`
    

    // 发货地址
    
    SendAddress   string `json:"send_address,omitempty"`
    

    // 订单总价(单位: 分)
    
    OrderAmount   string `json:"order_amount,omitempty"`
    

    // 猫超商品类目名称(从root到叶子节点)
    
    ItemCategory   string `json:"item_category,omitempty"`
    

    // 货物单价
    
    ItemPrice   string `json:"item_price,omitempty"`
    

    // 货物数量(下单销售数量)
    
    ItemQuantity   string `json:"item_quantity,omitempty"`
    

    // 货物名称
    
    ItemName   string `json:"item_name,omitempty"`
    

    // 交易子订单ID
    
    TbSubOrderId   int64 `json:"tb_sub_order_id,omitempty"`
    

}
*/

// InsuranceOrder 
type InsuranceOrder struct {

    // 签收时间
    SignTime   string `json:"sign_time,omitempty"`

    // 交易单号生成时间
    OrderCreateTime   string `json:"order_create_time,omitempty"`

    // 快递单号
    ExpressNo   string `json:"express_no,omitempty"`

    // 收货地址
    DeliveryAddress   string `json:"delivery_address,omitempty"`

    // 发货地址
    SendAddress   string `json:"send_address,omitempty"`

    // 订单总价(单位: 分)
    OrderAmount   string `json:"order_amount,omitempty"`

    // 猫超商品类目名称(从root到叶子节点)
    ItemCategory   string `json:"item_category,omitempty"`

    // 货物单价
    ItemPrice   string `json:"item_price,omitempty"`

    // 货物数量(下单销售数量)
    ItemQuantity   string `json:"item_quantity,omitempty"`

    // 货物名称
    ItemName   string `json:"item_name,omitempty"`

    // 交易子订单ID
    TbSubOrderId   int64 `json:"tb_sub_order_id,omitempty"`

}
