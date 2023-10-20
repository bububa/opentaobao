package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcaauthissueauthapplycallbackAPIRequest 出证申请回调 API请求
// alibaba.nazca.auth.issueauthapply.callback
//
// 出证申请回调
type AlibabanazcaauthissueauthapplycallbackAPIRequest struct {
	model.Params
	// 合同编号
	_contractNum string
	// 出证机构
	_issueOrg string
	// 客户在1688的唯一标识
	_platformUserId string
	// 出证报告下载地址
	_reportUrl string
	// 出证状态
	_status string
}

// NewAlibabanazcaauthissueauthapplycallbackRequest 初始化AlibabanazcaauthissueauthapplycallbackAPIRequest对象
func NewAlibabanazcaauthissueauthapplycallbackRequest() *AlibabanazcaauthissueauthapplycallbackAPIRequest {
	return &AlibabanazcaauthissueauthapplycallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.auth.issueauthapply.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContractNum is ContractNum Setter
// 合同编号
func (r *AlibabanazcaauthissueauthapplycallbackAPIRequest) SetContractNum(_contractNum string) error {
	r._contractNum = _contractNum
	r.Set("contract_num", _contractNum)
	return nil
}

// GetContractNum ContractNum Getter
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetContractNum() string {
	return r._contractNum
}

// SetIssueOrg is IssueOrg Setter
// 出证机构
func (r *AlibabanazcaauthissueauthapplycallbackAPIRequest) SetIssueOrg(_issueOrg string) error {
	r._issueOrg = _issueOrg
	r.Set("issue_org", _issueOrg)
	return nil
}

// GetIssueOrg IssueOrg Getter
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetIssueOrg() string {
	return r._issueOrg
}

// SetPlatformUserId is PlatformUserId Setter
// 客户在1688的唯一标识
func (r *AlibabanazcaauthissueauthapplycallbackAPIRequest) SetPlatformUserId(_platformUserId string) error {
	r._platformUserId = _platformUserId
	r.Set("platform_user_id", _platformUserId)
	return nil
}

// GetPlatformUserId PlatformUserId Getter
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetPlatformUserId() string {
	return r._platformUserId
}

// SetReportUrl is ReportUrl Setter
// 出证报告下载地址
func (r *AlibabanazcaauthissueauthapplycallbackAPIRequest) SetReportUrl(_reportUrl string) error {
	r._reportUrl = _reportUrl
	r.Set("report_url", _reportUrl)
	return nil
}

// GetReportUrl ReportUrl Getter
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetReportUrl() string {
	return r._reportUrl
}

// SetStatus is Status Setter
// 出证状态
func (r *AlibabanazcaauthissueauthapplycallbackAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabanazcaauthissueauthapplycallbackAPIRequest) GetStatus() string {
	return r._status
}
