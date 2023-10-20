package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest 查询上游出库单明细（带追溯码信息） API请求
// alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode
//
// 查询上游出库单明细(带追溯码信息)
type AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 服务时间
	_licenseToken string
	// 单据号
	_billCode string
	// 发货企业renEntId
	_fromRefUserId string
	// 收货企业refEntId
	_toRefUserId string
	// 委托企业id
	_agentRefEntId string
}

// NewAlibabaalihealthdrugkytwesupbilldetailwithcodeRequest 初始化AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest对象
func NewAlibabaalihealthdrugkytwesupbilldetailwithcodeRequest() *AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest {
	return &AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetFromRefUserId is FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
	r._fromRefUserId = _fromRefUserId
	r.Set("from_ref_user_id", _fromRefUserId)
	return nil
}

// GetFromRefUserId FromRefUserId Getter
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetFromRefUserId() string {
	return r._fromRefUserId
}

// SetToRefUserId is ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) SetToRefUserId(_toRefUserId string) error {
	r._toRefUserId = _toRefUserId
	r.Set("to_ref_user_id", _toRefUserId)
	return nil
}

// GetToRefUserId ToRefUserId Getter
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetToRefUserId() string {
	return r._toRefUserId
}

// SetAgentRefEntId is AgentRefEntId Setter
// 委托企业id
func (r *AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}
