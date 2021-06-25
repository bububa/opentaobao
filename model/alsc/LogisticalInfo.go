package alsc

// LogisticalInfo 
type LogisticalInfo struct {

    // 送达时间
    ArriveTime   string `json:"arrive_time,omitempty"`

    // 物流状态：WAIT_DELIVERY("WAIT_DELIVERY", "待发货"),  WAIT_RECEIVE("WAIT_RECEIVE", "待收货"),  SUCCESS("SUCCESS", "确认收货"),  REFUND("REFUND", "退货");
    LogisticsStatus   string `json:"logistics_status,omitempty"`

    // 收款地址
    ReceiveAddress   string `json:"receive_address,omitempty"`

    // 联系人手机
    ReceivePhone   string `json:"receive_phone,omitempty"`

}
