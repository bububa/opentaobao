package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSearchbillAPIRequest 通过时间段批量查询入出库单信息 API请求
// alibaba.alihealth.drug.kyt.searchbill
//
// 通过时间段批量查询入出库单信息
type AlibabaAlihealthDrugKytSearchbillAPIRequest struct {
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
	// 当前页
	_curPage int64
	// 页大小
	_pageSize int64
	// 单据号码
	_billCode string
	// 单据类型  A : 所有  AI :入库    AO:出库
	_billType string
}

// NewAlibabaAlihealthDrugKytSearchbillRequest 初始化AlibabaAlihealthDrugKytSearchbillAPIRequest对象
func NewAlibabaAlihealthDrugKytSearchbillRequest() *AlibabaAlihealthDrugKytSearchbillAPIRequest {
	return &AlibabaAlihealthDrugKytSearchbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.searchbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntId is RefEntId Setter
// 企业标识
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetAuthRefUserId is AuthRefUserId Setter
// 货主
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetAuthRefUserId(_authRefUserId string) error {
	r._authRefUserId = _authRefUserId
	r.Set("auth_ref_user_id", _authRefUserId)
	return nil
}

// GetAuthRefUserId AuthRefUserId Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetAuthRefUserId() string {
	return r._authRefUserId
}

// SetBeginDate is BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPartnerIdSend is PartnerIdSend Setter
// 发货企业
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetPartnerIdSend(_partnerIdSend string) error {
	r._partnerIdSend = _partnerIdSend
	r.Set("partner_id_send", _partnerIdSend)
	return nil
}

// GetPartnerIdSend PartnerIdSend Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetPartnerIdSend() string {
	return r._partnerIdSend
}

// SetPartnerIdRecv is PartnerIdRecv Setter
// 收货企业
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetPartnerIdRecv(_partnerIdRecv string) error {
	r._partnerIdRecv = _partnerIdRecv
	r.Set("partner_id_recv", _partnerIdRecv)
	return nil
}

// GetPartnerIdRecv PartnerIdRecv Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetPartnerIdRecv() string {
	return r._partnerIdRecv
}

// SetAgentRefUserId is AgentRefUserId Setter
// 代理企业
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetCurPage is CurPage Setter
// 当前页
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetCurPage(_curPage int64) error {
	r._curPage = _curPage
	r.Set("cur_page", _curPage)
	return nil
}

// GetCurPage CurPage Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetCurPage() int64 {
	return r._curPage
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillType is BillType Setter
// 单据类型  A : 所有  AI :入库    AO:出库
func (r *AlibabaAlihealthDrugKytSearchbillAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytSearchbillAPIRequest) GetBillType() string {
	return r._billType
}
