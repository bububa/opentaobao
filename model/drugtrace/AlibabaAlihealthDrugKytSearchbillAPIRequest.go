package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytsearchbillAPIRequest 通过时间段批量查询入出库单信息 API请求
// alibaba.alihealth.drug.kyt.searchbill
//
// 通过时间段批量查询入出库单信息
type AlibabaalihealthdrugkytsearchbillAPIRequest struct {
	model.Params
	// 企业标识
	_refEntId string
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

// NewAlibabaalihealthdrugkytsearchbillRequest 初始化AlibabaalihealthdrugkytsearchbillAPIRequest对象
func NewAlibabaalihealthdrugkytsearchbillRequest() *AlibabaalihealthdrugkytsearchbillAPIRequest {
	return &AlibabaalihealthdrugkytsearchbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.searchbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业标识
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetAuthRefUserId is AuthRefUserId Setter
// 货主
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetAuthRefUserId(_authRefUserId string) error {
	r._authRefUserId = _authRefUserId
	r.Set("auth_ref_user_id", _authRefUserId)
	return nil
}

// GetAuthRefUserId AuthRefUserId Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetAuthRefUserId() string {
	return r._authRefUserId
}

// SetBeginDate is BeginDate Setter
// 开始日期
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPartnerIdSend is PartnerIdSend Setter
// 发货企业
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetPartnerIdSend(_partnerIdSend string) error {
	r._partnerIdSend = _partnerIdSend
	r.Set("partner_id_send", _partnerIdSend)
	return nil
}

// GetPartnerIdSend PartnerIdSend Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetPartnerIdSend() string {
	return r._partnerIdSend
}

// SetPartnerIdRecv is PartnerIdRecv Setter
// 收货企业
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetPartnerIdRecv(_partnerIdRecv string) error {
	r._partnerIdRecv = _partnerIdRecv
	r.Set("partner_id_recv", _partnerIdRecv)
	return nil
}

// GetPartnerIdRecv PartnerIdRecv Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetPartnerIdRecv() string {
	return r._partnerIdRecv
}

// SetAgentRefUserId is AgentRefUserId Setter
// 代理企业
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillType is BillType Setter
// 单据类型  A : 所有  AI :入库    AO:出库
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetBillType() string {
	return r._billType
}

// SetCurPage is CurPage Setter
// 当前页
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetCurPage(_curPage int64) error {
	r._curPage = _curPage
	r.Set("cur_page", _curPage)
	return nil
}

// GetCurPage CurPage Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetCurPage() int64 {
	return r._curPage
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaalihealthdrugkytsearchbillAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthdrugkytsearchbillAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
