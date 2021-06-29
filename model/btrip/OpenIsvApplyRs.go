package btrip

// OpenIsvApplyRs 
type OpenIsvApplyRs struct {
    // 商旅审批展示id
    ApplyShowId   string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
    // 审批人列表
    ApproverList   []OpenApproverInfo `json:"approver_list,omitempty" xml:"approver_list>open_approver_info,omitempty"`
    // 商旅企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 企业名称
    CorpName   string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
    // 申请人部门id
    DepartId   string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
    // 部门名称
    DepartName   string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 商旅审批单id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 行程列表
    ItineraryList   []OpenItineraryInfo `json:"itinerary_list,omitempty" xml:"itinerary_list>open_itinerary_info,omitempty"`
    // 审批状态：0审批中 1已同意 2已拒绝 3已转交，4已取消 5已终止 6发起审批 7评论
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 审批状态描述
    StatusDesc   string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
    // 第三方业务id
    ThirdpartBusinessId   string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
    // 第三方审批单id,如果非第三方审批单则为空
    ThirdpartId   string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
    // 出行人列表
    TravelerList   []OpenUserInfo `json:"traveler_list,omitempty" xml:"traveler_list>open_user_info,omitempty"`
    // 出差事由
    TripCause   string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
    // 出差天数
    TripDay   int64 `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
    // 审批单标题
    TripTitle   string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
    // 申请人id（第三方用户Id）
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 用户名称
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // 申请单提交类型 1代提交 2本人提交 注意：当申请单为代提交时，申请单提交人自己无法为自己下单
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 流程编码
    FlowCode   string `json:"flow_code,omitempty" xml:"flow_code,omitempty"`
}
