package filmtfavatar

// ReturnValue 
type ReturnValue struct {

    // 数据: 包含: 淘宝订单ID,系统商订单号,付款时间,核销时间,影院ID,影院名称,卖品名称,卖品内容,卖品来源,卖品结算单价,卖品数量,实际结算金额,卖品结算总价,影院卖品补贴,订单状态,退款状态,会员卡标识,备注,是否后结算,
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    

    // 数据数量
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

}
