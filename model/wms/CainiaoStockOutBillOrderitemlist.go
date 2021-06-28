package wms

// CainiaoStockOutBillOrderitemlist 
/* model for simplify = false
type CainiaoStockOutBillOrderitemlist struct {

    // 订单商品信息
    
    OrderItem  *struct {
        CainiaoStockOutBillOrderitem  *CainiaoStockOutBillOrderitem `json:"cainiao_stock_out_bill_orderitem,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// CainiaoStockOutBillOrderitemlist 
type CainiaoStockOutBillOrderitemlist struct {

    // 订单商品信息
    OrderItem   *CainiaoStockOutBillOrderitem `json:"order_item,omitempty"`

}
