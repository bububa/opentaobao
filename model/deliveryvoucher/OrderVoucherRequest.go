package deliveryvoucher

// OrderVoucherRequest 
type OrderVoucherRequest struct {
    // 操作时间
    OperateDate   string `json:"operate_date,omitempty" xml:"operate_date,omitempty"`
    // 扩展参数
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    // 流水号:唯一，幂等性判断
    OpId   string `json:"op_id,omitempty" xml:"op_id,omitempty"`
    // 第三方服务商标识：top appkey
    Provider   string `json:"provider,omitempty" xml:"provider,omitempty"`
    // 主订单id
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 券信息,券信息,最多100条券记录
    VoucherInfos   []DeliveryVoucherInfoDTO `json:"voucher_infos,omitempty" xml:"voucher_infos>delivery_voucher_info_dto,omitempty"`
    // 预约时间
    AppointmentTime   string `json:"appointment_time,omitempty" xml:"appointment_time,omitempty"`
    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 1:物流发货 2：自提
    OutOrderType   int64 `json:"out_order_type,omitempty" xml:"out_order_type,omitempty"`
    // 自提方式必传
    TerminalAddress   string `json:"terminal_address,omitempty" xml:"terminal_address,omitempty"`
    // 自提方式必传
    TerminalPhone   string `json:"terminal_phone,omitempty" xml:"terminal_phone,omitempty"`
    // 自提方式必传
    TerminalName   string `json:"terminal_name,omitempty" xml:"terminal_name,omitempty"`
}
