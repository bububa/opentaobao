package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytListpartsByagentAPIRequest 物流代货主查往来单位接口 API请求
// alibaba.alihealth.drug.kyt.listparts.byagent
//
// 代理企业查询往来单位列表
type AlibabaAlihealthDrugKytListpartsByagentAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytListpartsByagentRequest 初始化AlibabaAlihealthDrugKytListpartsByagentAPIRequest对象
func NewAlibabaAlihealthDrugKytListpartsByagentRequest() *AlibabaAlihealthDrugKytListpartsByagentAPIRequest {
	return &AlibabaAlihealthDrugKytListpartsByagentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.listparts.byagent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（货主企业）
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetAgentRefEntId is AgentRefEntId Setter
// 代理企业唯一标识（物流企业）
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListpartsByagentAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytListpartsByagentAPIRequest) GetPage() int64 {
	return r._page
}
