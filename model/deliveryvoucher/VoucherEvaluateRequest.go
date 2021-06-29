package deliveryvoucher

// VoucherEvaluateRequest 
type VoucherEvaluateRequest struct {
    // 评价内容
    EvaluateContent   string `json:"evaluate_content,omitempty" xml:"evaluate_content,omitempty"`
    // 评价时间
    EvaluateTime   string `json:"evaluate_time,omitempty" xml:"evaluate_time,omitempty"`
    // 主交易订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 评价分数 1：失望；2：不满；3：一般；4：满意；5：惊喜
    EvaluateScore   int64 `json:"evaluate_score,omitempty" xml:"evaluate_score,omitempty"`
    // 券信息，券信息,最多20条券记录
    VoucherInfos   []DeliveryVoucherInfoDTO `json:"voucher_infos,omitempty" xml:"voucher_infos>delivery_voucher_info_dto,omitempty"`
    // 可扩展字段
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    // 操作时间
    OperateDate   string `json:"operate_date,omitempty" xml:"operate_date,omitempty"`
    // 流水号:唯一，幂等性判断
    OpId   string `json:"op_id,omitempty" xml:"op_id,omitempty"`
    // 第三方服务商标识,top appkey
    Provider   string `json:"provider,omitempty" xml:"provider,omitempty"`
}
