package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest 物流代货主查找往来单位接口 API请求
// alibaba.alihealth.drug.kyt.wes.listparts.byagent
//
// 代理企业查询往来单位列表
type AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytWesListpartsByagentRequest 初始化AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest对象
func NewAlibabaAlihealthDrugKytWesListpartsByagentRequest() *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest {
	return &AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) Reset() {
	r._refEntId = ""
	r._licenseToken = ""
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
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.listparts.byagent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（货主企业）
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetAgentRefEntId is AgentRefEntId Setter
// 代理企业唯一标识（物流企业）
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) GetPage() int64 {
	return r._page
}

var poolAlibabaAlihealthDrugKytWesListpartsByagentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesListpartsByagentRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesListpartsByagentRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest
func GetAlibabaAlihealthDrugKytWesListpartsByagentAPIRequest() *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest {
	return poolAlibabaAlihealthDrugKytWesListpartsByagentAPIRequest.Get().(*AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesListpartsByagentAPIRequest 将 AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesListpartsByagentAPIRequest(v *AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesListpartsByagentAPIRequest.Put(v)
}
