package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSearchbillAPIRequest 通过时间段批量查询入出库单信息 API请求
// alibaba.alihealth.drug.kyt.wes.searchbill
//
// 通过时间段批量查询入出库单信息
type AlibabaAlihealthDrugKytWesSearchbillAPIRequest struct {
	model.Params
	// 企业标识
	_refEntId string
	// 服务时间
	_licenseToken string
	// 货主
	_authRefUserId string
	// 开始日期
	_beginDate string
	// 结束日期
	_endDate string
	// 发货企业
	_partnerIdSend string
	// 收货企业
	_partnerIdRecv string
	// 代理企业
	_agentRefUserId string
	// 单据号码
	_billCode string
	// 单据类型  A : 所有  AI :入库    AO:出库
	_billType string
	// 当前页
	_curPage int64
	// 页大小
	_pageSize int64
}

// NewAlibabaAlihealthDrugKytWesSearchbillRequest 初始化AlibabaAlihealthDrugKytWesSearchbillAPIRequest对象
func NewAlibabaAlihealthDrugKytWesSearchbillRequest() *AlibabaAlihealthDrugKytWesSearchbillAPIRequest {
	return &AlibabaAlihealthDrugKytWesSearchbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.searchbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业标识
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetAuthRefUserId is AuthRefUserId Setter
// 货主
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetAuthRefUserId(_authRefUserId string) error {
	r._authRefUserId = _authRefUserId
	r.Set("auth_ref_user_id", _authRefUserId)
	return nil
}

// GetAuthRefUserId AuthRefUserId Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetAuthRefUserId() string {
	return r._authRefUserId
}

// SetBeginDate is BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPartnerIdSend is PartnerIdSend Setter
// 发货企业
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetPartnerIdSend(_partnerIdSend string) error {
	r._partnerIdSend = _partnerIdSend
	r.Set("partner_id_send", _partnerIdSend)
	return nil
}

// GetPartnerIdSend PartnerIdSend Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetPartnerIdSend() string {
	return r._partnerIdSend
}

// SetPartnerIdRecv is PartnerIdRecv Setter
// 收货企业
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetPartnerIdRecv(_partnerIdRecv string) error {
	r._partnerIdRecv = _partnerIdRecv
	r.Set("partner_id_recv", _partnerIdRecv)
	return nil
}

// GetPartnerIdRecv PartnerIdRecv Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetPartnerIdRecv() string {
	return r._partnerIdRecv
}

// SetAgentRefUserId is AgentRefUserId Setter
// 代理企业
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillType is BillType Setter
// 单据类型  A : 所有  AI :入库    AO:出库
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetBillType() string {
	return r._billType
}

// SetCurPage is CurPage Setter
// 当前页
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetCurPage(_curPage int64) error {
	r._curPage = _curPage
	r.Set("cur_page", _curPage)
	return nil
}

// GetCurPage CurPage Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetCurPage() int64 {
	return r._curPage
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytWesSearchbillAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytWesSearchbillAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
