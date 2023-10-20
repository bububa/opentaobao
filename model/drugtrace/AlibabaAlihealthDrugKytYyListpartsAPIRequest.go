package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyListpartsAPIRequest 查询往来单位 API请求
// alibaba.alihealth.drug.kyt.yy.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytYyListpartsAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytYyListpartsRequest 初始化AlibabaAlihealthDrugKytYyListpartsAPIRequest对象
func NewAlibabaAlihealthDrugKytYyListpartsRequest() *AlibabaAlihealthDrugKytYyListpartsAPIRequest {
	return &AlibabaAlihealthDrugKytYyListpartsAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) Reset() {
	r._refEntId = ""
	r._entName = ""
	r._refPartnerId = ""
	r._beginDate = ""
	r._endDate = ""
	r._agentRefEntId = ""
	r._pageSize = 0
	r._page = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.listparts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（货主企业）
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetAgentRefEntId is AgentRefEntId Setter
// 代理企业唯一标识（物流企业）
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytYyListpartsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytYyListpartsAPIRequest) GetPage() int64 {
	return r._page
}

var poolAlibabaAlihealthDrugKytYyListpartsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytYyListpartsRequest()
	},
}

// GetAlibabaAlihealthDrugKytYyListpartsRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytYyListpartsAPIRequest
func GetAlibabaAlihealthDrugKytYyListpartsAPIRequest() *AlibabaAlihealthDrugKytYyListpartsAPIRequest {
	return poolAlibabaAlihealthDrugKytYyListpartsAPIRequest.Get().(*AlibabaAlihealthDrugKytYyListpartsAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytYyListpartsAPIRequest 将 AlibabaAlihealthDrugKytYyListpartsAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytYyListpartsAPIRequest(v *AlibabaAlihealthDrugKytYyListpartsAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytYyListpartsAPIRequest.Put(v)
}
