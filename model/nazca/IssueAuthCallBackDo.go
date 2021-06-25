package nazca

// IssueAuthCallBackDo 
type IssueAuthCallBackDo struct {

    // 合同编号
    ContractNum   string `json:"contract_num,omitempty"`

    // 出证机构
    IssueOrg   string `json:"issue_org,omitempty"`

    // 客户在1688的唯一标识
    PlatformUserId   string `json:"platform_user_id,omitempty"`

    // 出证报告下载地址
    ReportUrl   string `json:"report_url,omitempty"`

    // 出证状态 0：出证中；1：出证成功
    Status   string `json:"status,omitempty"`

}
