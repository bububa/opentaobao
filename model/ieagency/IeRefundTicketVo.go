package ieagency

// IeRefundTicketVo 
type IeRefundTicketVo struct {
    // 填写退款金额
    RefundMoney   int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
    // 申请单id
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 乘机人id
    PassengerId   int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
    // 申请时间
    ApplyTime   string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
    // 初始化:1, 退款失败2, 退款成功:3
    RefundPayStatus   int64 `json:"refund_pay_status,omitempty" xml:"refund_pay_status,omitempty"`
    // 票号
    TicketNo   string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
    // agentAgreeTime
    AgentAgreeTime   string `json:"agent_agree_time,omitempty" xml:"agent_agree_time,omitempty"`
    // 待处理:1, 已同意:2,已拒绝:3,已受理:6,已退款:7
    RefundStatus   int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    // 申请原因
    ApplyReason   string `json:"apply_reason,omitempty" xml:"apply_reason,omitempty"`
    // 代理商名称
    AgentName   string `json:"agent_name,omitempty" xml:"agent_name,omitempty"`
    // 实际退还买家
    RefundToUserMoney   int64 `json:"refund_to_user_money,omitempty" xml:"refund_to_user_money,omitempty"`
    // 乘机人名称
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    // agentPayTime
    AgentPayTime   string `json:"agent_pay_time,omitempty" xml:"agent_pay_time,omitempty"`
    // 自愿申请:0,非自愿退票(不可抗力):1,自愿退票（我要改变行程计划、我不想飞）:2,自愿退票（填错名字、选错日期、选错航班）:3,自愿退票（生病了无法乘机（无二甲医院证明））:4,非自愿退票（航班延误或取消、航班时刻变更等航司原因）:5,非自愿退票（身体原因且有二级甲等医院<含>以上的医院证明）:6,非自愿退票（旅客拒签或其他不可抗力因素）:7
    ApplyType   int64 `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
    // agentRefuseTime
    AgentRefuseTime   string `json:"agent_refuse_time,omitempty" xml:"agent_refuse_time,omitempty"`
    // 代理商回复
    ApplyAnswer   string `json:"apply_answer,omitempty" xml:"apply_answer,omitempty"`
    // 订单id
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // agentReceiveTime
    AgentReceiveTime   string `json:"agent_receive_time,omitempty" xml:"agent_receive_time,omitempty"`
    // 2^0=1 限时免费退；2^1=2 极速退款；比如表示(限时免费退+实时退款)=1+2=3
    RefundProductType   int64 `json:"refund_product_type,omitempty" xml:"refund_product_type,omitempty"`
    // 是否补退订单
    BuTui   bool `json:"bu_tui,omitempty" xml:"bu_tui,omitempty"`
}
