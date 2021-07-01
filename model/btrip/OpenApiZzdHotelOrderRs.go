package btrip

// OpenApiZzdHotelOrderRs 
type OpenApiZzdHotelOrderRs struct {
    // 订单状态：1, "等待确认";2, "等待付款"; 3, "预订成功";4, "申请退款";5, "退款成功";6, "已关闭";0,"点击详情查看状态";7, "结账成功"; 8, "支付成功";
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 服务费（分）
    ServiceFee   int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
    // 退订服务费（分）
    RefundServiceFee   int64 `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
    // 退票费（分）
    RefundFee   int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 个人支付（分），当pay_type是个人支付或混合支付时有值
    PersonalAmount   int64 `json:"personal_amount,omitempty" xml:"personal_amount,omitempty"`
    // 企业支付（分），当pay_type是企业支付或混合支付时有值
    CorpAmount   int64 `json:"corp_amount,omitempty" xml:"corp_amount,omitempty"`
    // 总结算金额（分）
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    // 结算方式:1：个人现付，2:企业现付,4:企业月结，8、企业预存
    SettleType   int64 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
    // 支付类型,0:企业支付，1:个人支付,2.混合支付
    PayType   int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    // 订单状态改变时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 申请单标题
    ApplyTitle   string `json:"apply_title,omitempty" xml:"apply_title,omitempty"`
    // 第三方申请单id
    ThirdPartApplyId   string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
    // 商旅审批单id
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 部门名称
    DepartName   string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
    // 用户名称
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    // 用户ID
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 第三方企业ID
    ThirdpartCorpId   string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
    // 商旅企业ID
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 资金流向:1:支出，2:收入
    TradeAction   int64 `json:"trade_action,omitempty" xml:"trade_action,omitempty"`
    // 交易类型：1:酒店费用，2,："酒店订单服务费"，101:酒店退款'，102,："酒店赔付"
    TradeType   int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    // 第三方交易ID
    TradeId   string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
    // 货币种类
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 审批人名称，多个/分开
    Approver   string `json:"approver,omitempty" xml:"approver,omitempty"`
    // 结算批次（该订单结算时间）
    SettlementBatchNo   string `json:"settlement_batch_no,omitempty" xml:"settlement_batch_no,omitempty"`
    // 差标
    TravelStandardAmount   int64 `json:"travel_standard_amount,omitempty" xml:"travel_standard_amount,omitempty"`
    // 商旅发票ID
    InvoiceId   string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
    // 发票抬头
    InvoiceTitle   string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
    // 项目编号
    ProjectCode   string `json:"project_code,omitempty" xml:"project_code,omitempty"`
    // 项目名称
    ProjectTitle   string `json:"project_title,omitempty" xml:"project_title,omitempty"`
    // 第三方行程单号
    ThirdpartItineraryId   string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
    // 商旅行程单号
    ItineraryNo   string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
    // 成本中心信息
    CostCenterList   []CostCenterDo `json:"cost_center_list,omitempty" xml:"cost_center_list>cost_center_do,omitempty"`
    // 用户信息
    ClientInfoDos   []ClientInfoDo `json:"client_info_dos,omitempty" xml:"client_info_dos>client_info_do,omitempty"`
    // 订单类型：1：全额支付、5：到店支付、6：信用住',
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 入住时间
    CheckInDate   string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
    // 离店时间
    CheckOutDate   string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
    // 酒店信息
    HotelInfoDo   *HotelInfoDo `json:"hotel_info_do,omitempty" xml:"hotel_info_do,omitempty"`
    // 房间信息
    RoomInfoDo   *RoomInfoDo `json:"room_info_do,omitempty" xml:"room_info_do,omitempty"`
}
