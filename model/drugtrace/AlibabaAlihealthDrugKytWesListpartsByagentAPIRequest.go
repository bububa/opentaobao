package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest 物流代货主查找往来单位接口 API请求
// alibaba.alihealth.drug.kyt.wes.listparts.byagent
//
// 代理企业查询往来单位列表
type AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest struct {
	model.Params
	// 企业唯一标识（货主企业）
	_refEntId string
	// 服务时间
	_licenseToken string
	// 企业名称
	_entName string
	// 企业自定义编号
	_refPartnerId string
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
	// 代理企业唯一标识（物流企业）
	_agentRefEntId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaalihealthdrugkytweslistpartsbyagentRequest 初始化AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest对象
func NewAlibabaalihealthdrugkytweslistpartsbyagentRequest() *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest {
	return &AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.listparts.byagent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（货主企业）
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetAgentRefEntId is AgentRefEntId Setter
// 代理企业唯一标识（物流企业）
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest) GetPage() int64 {
	return r._page
}
