package wms

// CainiaoStockInBillOrderitemlist 
/* model for simplify = false
type CainiaoStockInBillOrderitemlist struct {

    // 入库单信息
    
    OrderItem  *struct {
        CainiaoStockInBillOrderitem  *CainiaoStockInBillOrderitem `json:"cainiao_stock_in_bill_orderitem,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// CainiaoStockInBillOrderitemlist 
type CainiaoStockInBillOrderitemlist struct {

    // 入库单信息
    OrderItem   *CainiaoStockInBillOrderitem `json:"order_item,omitempty"`

}
