package wms

// CainiaoReturnBillOrderitemlist 
/* model for simplify = false
type CainiaoReturnBillOrderitemlist struct {

    // 订单商品信息
    
    OrderItem  *struct {
        CainiaoReturnBillOrderitem  *CainiaoReturnBillOrderitem `json:"cainiao_return_bill_orderitem,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// CainiaoReturnBillOrderitemlist 
type CainiaoReturnBillOrderitemlist struct {

    // 订单商品信息
    OrderItem   *CainiaoReturnBillOrderitem `json:"order_item,omitempty"`

}
