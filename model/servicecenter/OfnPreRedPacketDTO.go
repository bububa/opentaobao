package servicecenter

// OfnPreRedPacketDTO 
type OfnPreRedPacketDTO struct {

    // 操作列表
    
    ActionList   []OfnPreRedPacketActionDto `json:"action_list,omitempty" xml:"action_list,omitempty"`
    

    // 活动id
    
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    

    // 是否信用单
    
    CreditPay   bool `json:"credit_pay,omitempty" xml:"credit_pay,omitempty"`
    

    // 信用额度，单位分
    
    CreditPayLimit   int64 `json:"credit_pay_limit,omitempty" xml:"credit_pay_limit,omitempty"`
    

    // 尾款部分-已经发放的金额，单位分
    
    EndAlreadySendAmount   int64 `json:"end_already_send_amount,omitempty" xml:"end_already_send_amount,omitempty"`
    

    // 尾款部分-计划发放的金额，单位分
    
    EndPlanSendAmount   int64 `json:"end_plan_send_amount,omitempty" xml:"end_plan_send_amount,omitempty"`
    

    // 尾款部分-等待发放的金额，单位分
    
    EndWaitSendAmount   int64 `json:"end_wait_send_amount,omitempty" xml:"end_wait_send_amount,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 主键
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 新机id
    
    NewItemId   int64 `json:"new_item_id,omitempty" xml:"new_item_id,omitempty"`
    

    // 新机订单id
    
    NewOrderId   int64 `json:"new_order_id,omitempty" xml:"new_order_id,omitempty"`
    

    // 旧机质检价
    
    OldItemActualPrice   int64 `json:"old_item_actual_price,omitempty" xml:"old_item_actual_price,omitempty"`
    

    // 旧机评估价
    
    OldItemApprizePrice   int64 `json:"old_item_apprize_price,omitempty" xml:"old_item_apprize_price,omitempty"`
    

    // 旧机id
    
    OldItemSpuId   int64 `json:"old_item_spu_id,omitempty" xml:"old_item_spu_id,omitempty"`
    

    // 旧机单id
    
    OldOrderId   int64 `json:"old_order_id,omitempty" xml:"old_order_id,omitempty"`
    

    // 计划发放的金额，单位分
    
    PlanSendAmount   int64 `json:"plan_send_amount,omitempty" xml:"plan_send_amount,omitempty"`
    

    // 预发部分-已经发放的金额，单位分
    
    PreAlreadySendAmount   int64 `json:"pre_already_send_amount,omitempty" xml:"pre_already_send_amount,omitempty"`
    

    // 预发部分-计划发放的金额，单位分
    
    PrePlanSendAmount   int64 `json:"pre_plan_send_amount,omitempty" xml:"pre_plan_send_amount,omitempty"`
    

    // 预发部分-等待发放的金额，单位分
    
    PreWaitSendAmount   int64 `json:"pre_wait_send_amount,omitempty" xml:"pre_wait_send_amount,omitempty"`
    

    // 状态。初始化=1，重试中=2，失败=3，成功=4
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 由回收商已经扣回的金额，单位分
    
    TmallAlreadyDeductAmount   int64 `json:"tmall_already_deduct_amount,omitempty" xml:"tmall_already_deduct_amount,omitempty"`
    

    // 等待回收商扣回的金额，单位分
    
    TmallWaitDeductAmount   int64 `json:"tmall_wait_deduct_amount,omitempty" xml:"tmall_wait_deduct_amount,omitempty"`
    

    // 由天猫已经扣回的金额，单位分
    
    TpAlreadyDeductAmount   int64 `json:"tp_already_deduct_amount,omitempty" xml:"tp_already_deduct_amount,omitempty"`
    

    // 等待天猫扣回的金额，单位分
    
    TpWaitDeductAmount   int64 `json:"tp_wait_deduct_amount,omitempty" xml:"tp_wait_deduct_amount,omitempty"`
    

    // 乐观锁
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

    // 新机优惠价
    
    NewItemCouponPrice   int64 `json:"new_item_coupon_price,omitempty" xml:"new_item_coupon_price,omitempty"`
    

}
