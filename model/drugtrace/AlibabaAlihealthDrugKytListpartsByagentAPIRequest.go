package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytlistpartsbyagentAPIRequest 物流代货主查往来单位接口 API请求
// alibaba.alihealth.drug.kyt.listparts.byagent
//
// 代理企业查询往来单位列表
type AlibabaalihealthdrugkytlistpartsbyagentAPIRequest struct {
	model.Params
	// 企业唯一标识（货主企业）
	_refEntId string
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

// NewAlibabaalihealthdrugkytlistpartsbyagentRequest 初始化AlibabaalihealthdrugkytlistpartsbyagentAPIRequest对象
func NewAlibabaalihealthdrugkytlistpartsbyagentRequest() *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest {
	return &AlibabaalihealthdrugkytlistpartsbyagentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.listparts.byagent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（货主企业）
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetAgentRefEntId is AgentRefEntId Setter
// 代理企业唯一标识（物流企业）
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthdrugkytlistpartsbyagentAPIRequest) GetPage() int64 {
	return r._page
}
