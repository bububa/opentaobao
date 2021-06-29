package btrip

// OpenApplyRs 
type OpenApplyRs struct {
    // 审批单状态：0申请 1同意 2拒绝 3转交 4取消 5修改已同意 6撤销已同意 7修改审批中 8已同意(修改被拒绝) 9撤销审批中 10已同意(撤销被拒绝) 11已同意(修改被取消) 12已同意(撤销被取消)
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 出差标题
    TripTitle   string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
    // 出差理由
    TripCause   string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
    // 出差天数
    TripDay   int64 `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
    // 部门名称
    DepartName   string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
    // 部门id
    DepartId   string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
    // 用户名称
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // 用户id
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 企业名称
    CorpName   string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
    // 企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 第三方审批单id, 非第三方审批单则为空
    ThirdpartId   string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 审批单id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 行程单列表
    ItineraryList   []OpenItineraryInfo `json:"itinerary_list,omitempty" xml:"itinerary_list>open_itinerary_info,omitempty"`
    // 出行人列表
    TravelerList   []OpenUserInfo `json:"traveler_list,omitempty" xml:"traveler_list>open_user_info,omitempty"`
    // 审批人列表
    ApproverList   []OpenApproverInfo `json:"approver_list,omitempty" xml:"approver_list>open_approver_info,omitempty"`
    // 审批单状态描述
    StatusDesc   string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
    // 审批单展示id
    ApplyShowId   string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
}
