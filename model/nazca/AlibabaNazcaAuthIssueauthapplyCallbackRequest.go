package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出证申请回调 API请求
alibaba.nazca.auth.issueauthapply.callback

出证申请回调
*/
type AlibabaNazcaAuthIssueauthapplyCallbackRequest struct {
    model.Params
    // 合同编号
    contractNum   string
    // 出证机构
    issueOrg   string
    // 客户在1688的唯一标识
    platformUserId   string
    // 出证报告下载地址
    reportUrl   string
    // 出证状态
    status   string
}

// 初始化AlibabaNazcaAuthIssueauthapplyCallbackRequest对象
func NewAlibabaNazcaAuthIssueauthapplyCallbackRequest() *AlibabaNazcaAuthIssueauthapplyCallbackRequest{
    return &AlibabaNazcaAuthIssueauthapplyCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetApiMethodName() string {
    return "alibaba.nazca.auth.issueauthapply.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContractNum Setter
// 合同编号
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetContractNum(contractNum string) error {
    r.contractNum = contractNum
    r.Set("contract_num", contractNum)
    return nil
}

// ContractNum Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetContractNum() string {
    return r.contractNum
}
// IssueOrg Setter
// 出证机构
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetIssueOrg(issueOrg string) error {
    r.issueOrg = issueOrg
    r.Set("issue_org", issueOrg)
    return nil
}

// IssueOrg Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetIssueOrg() string {
    return r.issueOrg
}
// PlatformUserId Setter
// 客户在1688的唯一标识
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetPlatformUserId(platformUserId string) error {
    r.platformUserId = platformUserId
    r.Set("platform_user_id", platformUserId)
    return nil
}

// PlatformUserId Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetPlatformUserId() string {
    return r.platformUserId
}
// ReportUrl Setter
// 出证报告下载地址
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetReportUrl(reportUrl string) error {
    r.reportUrl = reportUrl
    r.Set("report_url", reportUrl)
    return nil
}

// ReportUrl Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetReportUrl() string {
    return r.reportUrl
}
// Status Setter
// 出证状态
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetStatus() string {
    return r.status
}
