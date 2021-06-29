package xhotelonlineorder

// InvoiceInfoParam 
type InvoiceInfoParam struct {
    // 请求id 唯一值(同极速开票开票请求回传request_id值)
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 酒店开票点税号
    TaxNum   string `json:"tax_num,omitempty" xml:"tax_num,omitempty"`
    // 实际房费(分)
    InvoiceRoomPrice   int64 `json:"invoice_room_price,omitempty" xml:"invoice_room_price,omitempty"`
    // 实际杂费(分)
    InvoiceOtherPrice   int64 `json:"invoice_other_price,omitempty" xml:"invoice_other_price,omitempty"`
    // 发票类型（1:普通发票；2：增值税专用发票）
    InvoiceType   int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
    // 专票信息
    ValueAddedInfo   *ValueAddedInfo `json:"value_added_info,omitempty" xml:"value_added_info,omitempty"`
    // 领取方式（      0：前台自取  3: 送至房间）
    PostType   int64 `json:"post_type,omitempty" xml:"post_type,omitempty"`
    // 身份证后4位
    ShortIdNumber   string `json:"short_id_number,omitempty" xml:"short_id_number,omitempty"`
    // 预计开票房费金额(分)
    PlanInvoiceRoomPrice   int64 `json:"plan_invoice_room_price,omitempty" xml:"plan_invoice_room_price,omitempty"`
    // 公司抬头
    CompanyTitle   string `json:"company_title,omitempty" xml:"company_title,omitempty"`
    // 房间号
    RoomNum   string `json:"room_num,omitempty" xml:"room_num,omitempty"`
    // 需要发票时间(格式yyyy-MM-dd HH:mm:ss)
    WantTime   string `json:"want_time,omitempty" xml:"want_time,omitempty"`
    // 预计开票杂费金额(分)
    PlanInvoiceOtherPrice   int64 `json:"plan_invoice_other_price,omitempty" xml:"plan_invoice_other_price,omitempty"`
    // 用户渠道(0:未知,1:淘宝)
    UserChannel   int64 `json:"user_channel,omitempty" xml:"user_channel,omitempty"`
    // 酒店外部订单号(从查询pms账单接口中获取)
    OutOrderNum   string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
    // 用户名
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
}
