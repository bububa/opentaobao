package waybill

// WaybillDetailQueryRequest 
type WaybillDetailQueryRequest struct {

    // 电子面单单号
    WaybillCodes   []String `json:"waybill_codes,omitempty"`

    // CP快递公司编码
    CpCode   string `json:"cp_code,omitempty"`

    // 0:根据cp_code和waybil_code查询;1:根据订单号查询
    QueryBy   int64 `json:"query_by,omitempty"`

    // 交易订单号
    TradeOrderList   []String `json:"trade_order_list,omitempty"`

}
