package maitix

// MoaOrderParam 
type MoaOrderParam struct {
    // 项目ID-必填
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 场次ID-必填
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    // 外部订单号，即分销方订单号，最长不超过32位-必填
    ThirdOrderNo   string `json:"third_order_no,omitempty" xml:"third_order_no,omitempty"`
    // 订单总价-必填
    TotalPrice   int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
    // 必填-支付金额
    Payment   int64 `json:"payment,omitempty" xml:"payment,omitempty"`
    // 必填-出票方式 1=纸质票 2=身份证电子票 3=二维码电子票 4=短信电子票
    TicketMode   int64 `json:"ticket_mode,omitempty" xml:"ticket_mode,omitempty"`
    // 必填-购票类型 1=自助选座 2=系统购买 非选座类，全部为2
    BuyType   int64 `json:"buy_type,omitempty" xml:"buy_type,omitempty"`
    // 配送地址,有条件必填,如果是快递票则必填
    DeliverAddress   string `json:"deliver_address,omitempty" xml:"deliver_address,omitempty"`
    // 必填-取票方式： 1，无纸化；2，快递票；3，自助换票；4，门店自取。1和3为电子票，2和4为纸质票。
    DeliveryType   int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
    // 必填-操作人登录账号，可以在大麦分销后台查询
    OperatorLoginId   string `json:"operator_login_id,omitempty" xml:"operator_login_id,omitempty"`
    // 必填，随便填一个就行,支付方式 1：现金，2：支付宝，3：微信，5：银行卡
    PayType   int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    // 必填-购票人姓名
    RealTicketBuyerName   string `json:"real_ticket_buyer_name,omitempty" xml:"real_ticket_buyer_name,omitempty"`
    // 购票人证件号,如果是一单一证的项目必填
    RealTicketBuyerIdCardNo   string `json:"real_ticket_buyer_id_card_no,omitempty" xml:"real_ticket_buyer_id_card_no,omitempty"`
    // 购票人证件类型，1：身份证，2：护照；3：港澳通行证，4：台胞证，5：士兵/军官-如果是一单一证的项目必填
    RealTicketBuyerIdCardType   int64 `json:"real_ticket_buyer_id_card_type,omitempty" xml:"real_ticket_buyer_id_card_type,omitempty"`
    // 必填-购票人电话
    RealTicketBuyerPhone   string `json:"real_ticket_buyer_phone,omitempty" xml:"real_ticket_buyer_phone,omitempty"`
    // 必填-购票人电话国家代码
    RealTicketBuyerPhoneCountryCode   string `json:"real_ticket_buyer_phone_country_code,omitempty" xml:"real_ticket_buyer_phone_country_code,omitempty"`
    // 必填-是否系统自动选座，无座项目都是true.有座项目通过h5选座就填false.否则true
    AutoSelectSeats   bool `json:"auto_select_seats,omitempty" xml:"auto_select_seats,omitempty"`
    // 座位参数，必填，一张票就要有一个这个对象,如果一个套票由两张票组成。这个也就是2个对象,票品id是套票的id.
    SeatProps   []MoaTicketInfo `json:"seat_props,omitempty" xml:"seat_props>moa_ticket_info,omitempty"`
    // 下单票品信息
    TicketItems   []MoaTicketItemSpec `json:"ticket_items,omitempty" xml:"ticket_items>moa_ticket_item_spec,omitempty"`
    // 联系人-必填
    ContactInfo   *MoaOrderContactInfo `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
    // 备注-可选
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    // 超时取消时间,单位分种
    TimeoutMinutes   int64 `json:"timeout_minutes,omitempty" xml:"timeout_minutes,omitempty"`
    // 运费-如果是快递票，必填
    FarePrice   int64 `json:"fare_price,omitempty" xml:"fare_price,omitempty"`
    // 区域地址-如果是快递票必填
    AddressInfo   *MoaAddressInfo `json:"address_info,omitempty" xml:"address_info,omitempty"`
}
