package waybill

// WaybillDetailQueryRequest 
/* model for simplify = false
type WaybillDetailQueryRequest struct {

    // 电子面单单号
    
    WaybillCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"waybill_codes,omitempty"`
    

    // CP快递公司编码
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 0:根据cp_code和waybil_code查询;1:根据订单号查询
    
    QueryBy   int64 `json:"query_by,omitempty"`
    

    // 交易订单号
    
    TradeOrderList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"trade_order_list,omitempty"`
    

}
*/

// WaybillDetailQueryRequest 
type WaybillDetailQueryRequest struct {

    // 电子面单单号
    WaybillCodes   []string `json:"waybill_codes,omitempty"`

    // CP快递公司编码
    CpCode   string `json:"cp_code,omitempty"`

    // 0:根据cp_code和waybil_code查询;1:根据订单号查询
    QueryBy   int64 `json:"query_by,omitempty"`

    // 交易订单号
    TradeOrderList   []string `json:"trade_order_list,omitempty"`

}
