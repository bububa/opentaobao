package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqydelbillinfoAPIRequest 根据单据号删除单据 API请求
// alibaba.alihealth.drug.kyt.scqy.delbillinfo
//
// 根据单据号删除单据
type AlibabaalihealthdrugkytscqydelbillinfoAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 单据号
	_billCode string
	// 删除原因
	_deleteReason string
	// 删除人手机号
	_mobilePhone string
	// 当单据是代理企业上传的，需要填写代理企业的refEntId
	_ownerRefUserId string
}

// NewAlibabaalihealthdrugkytscqydelbillinfoRequest 初始化AlibabaalihealthdrugkytscqydelbillinfoAPIRequest对象
func NewAlibabaalihealthdrugkytscqydelbillinfoRequest() *AlibabaalihealthdrugkytscqydelbillinfoAPIRequest {
	return &AlibabaalihealthdrugkytscqydelbillinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.delbillinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetDeleteReason is DeleteReason Setter
// 删除原因
func (r *AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) SetDeleteReason(_deleteReason string) error {
	r._deleteReason = _deleteReason
	r.Set("delete_reason", _deleteReason)
	return nil
}

// GetDeleteReason DeleteReason Getter
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetDeleteReason() string {
	return r._deleteReason
}

// SetMobilePhone is MobilePhone Setter
// 删除人手机号
func (r *AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) SetMobilePhone(_mobilePhone string) error {
	r._mobilePhone = _mobilePhone
	r.Set("mobile_phone", _mobilePhone)
	return nil
}

// GetMobilePhone MobilePhone Getter
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetMobilePhone() string {
	return r._mobilePhone
}

// SetOwnerRefUserId is OwnerRefUserId Setter
// 当单据是代理企业上传的，需要填写代理企业的refEntId
func (r *AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) SetOwnerRefUserId(_ownerRefUserId string) error {
	r._ownerRefUserId = _ownerRefUserId
	r.Set("owner_ref_user_id", _ownerRefUserId)
	return nil
}

// GetOwnerRefUserId OwnerRefUserId Getter
func (r AlibabaalihealthdrugkytscqydelbillinfoAPIRequest) GetOwnerRefUserId() string {
	return r._ownerRefUserId
}
