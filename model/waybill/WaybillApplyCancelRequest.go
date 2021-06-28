package waybill

// WaybillApplyCancelRequest 
type WaybillApplyCancelRequest struct {

    // 面单使用者编号
    
    RealUserId   int64 `json:"real_user_id,omitempty" xml:"real_user_id,omitempty"`
    

    // 交易订单列表
    
    TradeOrderList   []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
    

    // CP快递公司编码
    
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    

    // 电子面单号码
    
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
    

    // ERP订单号或包裹号
    
    PackageId   string `json:"package_id,omitempty" xml:"package_id,omitempty"`
    

}
