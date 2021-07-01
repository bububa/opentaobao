package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytYyListpartsAPIRequest
查询往来单位 API请求
alibaba.alihealth.drug.kyt.yy.listparts

查询往来单位列表 */
type AlibabaAlihealthDrugKytYyListpartsAPIRequest struct {
	model.Params
	// 企业唯一标识（货主企业）
	_refEntId string
	// 企业名称
	_entName string
	// 企业自定义编号
	_refPartnerId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
	// 代理企业唯一标识（物流企业）
	_agentRefEntId string
}

// NewAlibabaAlihealthDrugKytYyListpartsRequest 初始化AlibabaAlihealthDrugKytYyListpartsAPIRequest对象
func NewAlibabaAlihealthDrugKytYyListpartsRequest() *AlibabaAlihealthDrugKytYyListpartsAPIRequest {
	return &AlibabaAlihealthDrugKytYyListpartsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.listparts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业唯一标识（货主企业）
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// Get EntName Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetEntName() string {
	return r._entName
}

// Set is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// Get RefPartnerId Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// Set is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetPage() int64 {
	return r._page
}

// Set is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// Get BeginDate Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// Set is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is AgentRefEntId Setter
// 代理企业唯一标识（物流企业）
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// Get AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}
