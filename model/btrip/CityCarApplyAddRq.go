package btrip

// CityCarApplyAddRq 
type CityCarApplyAddRq struct {
    // 出差事由
    Cause   string `json:"cause,omitempty" xml:"cause,omitempty"`
    // 用车城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 第三方企业ID
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 用车时间，按天管控，比如传值2021-03-18 20:26:56表示2021-03-18当天可用车，跨天情况配合finished_date参数使用
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    // 审批单关联的项目code
    ProjectCode   string `json:"project_code,omitempty" xml:"project_code,omitempty"`
    // 审批单关联的项目名
    ProjectName   string `json:"project_name,omitempty" xml:"project_name,omitempty"`
    // 审批单状态：0-申请，1-同意，2-拒绝
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 三方审批单ID
    ThirdPartApplyId   string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
    // 审批单关联的三方成本中心ID
    ThirdPartCostCenterId   string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
    // 审批单关联的三方发票抬头ID
    ThirdPartInvoiceId   string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
    // 审批单可用总次数
    TimesTotal   int64 `json:"times_total,omitempty" xml:"times_total,omitempty"`
    // 审批单可用次数类型：1-次数不限制，2-用户可指定次数，3-管理员限制次数；如果企业没有限制审批单使用次数的需求，这个参数传1(次数不限制)，同时times_total和times_used都传0即可
    TimesType   int64 `json:"times_type,omitempty" xml:"times_type,omitempty"`
    // 审批单已用次数
    TimesUsed   int64 `json:"times_used,omitempty" xml:"times_used,omitempty"`
    // 审批单标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 发起审批的第三方员工ID
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 用车截止时间，按天管控，比如date传值2021-03-18 20:26:56、finished_date传值2021-03-30 20:26:56表示2021-03-18(含)到2021-03-30(含)之间可用车，该参数不传值情况使用date作为用车截止时间；
    FinishedDate   string `json:"finished_date,omitempty" xml:"finished_date,omitempty"`
}
