package wms

// CainiaoReturnBillReturnorderinfo 
/* model for simplify = false
type CainiaoReturnBillReturnorderinfo struct {

    // ERP订单号
    
    OrderCode   string `json:"order_code,omitempty"`
    

    // 仓库订单编码，LBX号
    
    CnOrderCode   string `json:"cn_order_code,omitempty"`
    

    // 单据类型： 501 退货入库
    
    OrderType   int64 `json:"order_type,omitempty"`
    

    // 仓库订单完成时间
    
    ConfirmTime   string `json:"confirm_time,omitempty"`
    

    // 订单商品信息列表
    
    OrderItemList  struct {
        CainiaoReturnBillOrderitemlist  []CainiaoReturnBillOrderitemlist `json:"cainiao_return_bill_orderitemlist,omitempty"`
    } `json:"order_item_list,omitempty"`
    

    // 此销退订单对应原销售订单的菜鸟订单号
    
    PreCnOrderCode   string `json:"pre_cn_order_code,omitempty"`
    

}
*/

// CainiaoReturnBillReturnorderinfo 
type CainiaoReturnBillReturnorderinfo struct {

    // ERP订单号
    OrderCode   string `json:"order_code,omitempty"`

    // 仓库订单编码，LBX号
    CnOrderCode   string `json:"cn_order_code,omitempty"`

    // 单据类型： 501 退货入库
    OrderType   int64 `json:"order_type,omitempty"`

    // 仓库订单完成时间
    ConfirmTime   string `json:"confirm_time,omitempty"`

    // 订单商品信息列表
    OrderItemList   []CainiaoReturnBillOrderitemlist `json:"order_item_list,omitempty"`

    // 此销退订单对应原销售订单的菜鸟订单号
    PreCnOrderCode   string `json:"pre_cn_order_code,omitempty"`

}
