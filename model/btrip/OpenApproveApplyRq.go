package btrip

// OpenApproveApplyRq 
type OpenApproveApplyRq struct {

    // 审批单id
    
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    

    // 操作时间
    
    OperateTime   string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
    

    // 1已同意 2已拒绝 3已转交 4已取消
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 审批批注
    
    Note   string `json:"note,omitempty" xml:"note,omitempty"`
    

    // 企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 审批人姓名
    
    ApprovalUserName   string `json:"approval_user_name,omitempty" xml:"approval_user_name,omitempty"`
    

    // 审批人id
    
    ApprovalUserId   string `json:"approval_user_id,omitempty" xml:"approval_user_id,omitempty"`
    

}
