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
    _contractNum   string
    // 出证机构
    _issueOrg   string
    // 客户在1688的唯一标识
    _platformUserId   string
    // 出证报告下载地址
    _reportUrl   string
    // 出证状态
    _status   string
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
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetContractNum(_contractNum string) error {
    r._contractNum = _contractNum
    r.Set("contract_num", _contractNum)
    return nil
}

// ContractNum Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetContractNum() string {
    return r._contractNum
}
// IssueOrg Setter
// 出证机构
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetIssueOrg(_issueOrg string) error {
    r._issueOrg = _issueOrg
    r.Set("issue_org", _issueOrg)
    return nil
}

// IssueOrg Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetIssueOrg() string {
    return r._issueOrg
}
// PlatformUserId Setter
// 客户在1688的唯一标识
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetPlatformUserId(_platformUserId string) error {
    r._platformUserId = _platformUserId
    r.Set("platform_user_id", _platformUserId)
    return nil
}

// PlatformUserId Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetPlatformUserId() string {
    return r._platformUserId
}
// ReportUrl Setter
// 出证报告下载地址
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetReportUrl(_reportUrl string) error {
    r._reportUrl = _reportUrl
    r.Set("report_url", _reportUrl)
    return nil
}

// ReportUrl Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetReportUrl() string {
    return r._reportUrl
}
// Status Setter
// 出证状态
func (r *AlibabaNazcaAuthIssueauthapplyCallbackRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaNazcaAuthIssueauthapplyCallbackRequest) GetStatus() string {
    return r._status
}
