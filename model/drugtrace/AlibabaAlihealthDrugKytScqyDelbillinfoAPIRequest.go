package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest 根据单据号删除单据 API请求
// alibaba.alihealth.drug.kyt.scqy.delbillinfo
//
// 根据单据号删除单据
type AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytScqyDelbillinfoRequest 初始化AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytScqyDelbillinfoRequest() *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest {
	return &AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) Reset() {
	r._refEntId = ""
	r._billCode = ""
	r._deleteReason = ""
	r._mobilePhone = ""
	r._ownerRefUserId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.delbillinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetDeleteReason is DeleteReason Setter
// 删除原因
func (r *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) SetDeleteReason(_deleteReason string) error {
	r._deleteReason = _deleteReason
	r.Set("delete_reason", _deleteReason)
	return nil
}

// GetDeleteReason DeleteReason Getter
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetDeleteReason() string {
	return r._deleteReason
}

// SetMobilePhone is MobilePhone Setter
// 删除人手机号
func (r *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) SetMobilePhone(_mobilePhone string) error {
	r._mobilePhone = _mobilePhone
	r.Set("mobile_phone", _mobilePhone)
	return nil
}

// GetMobilePhone MobilePhone Getter
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetMobilePhone() string {
	return r._mobilePhone
}

// SetOwnerRefUserId is OwnerRefUserId Setter
// 当单据是代理企业上传的，需要填写代理企业的refEntId
func (r *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) SetOwnerRefUserId(_ownerRefUserId string) error {
	r._ownerRefUserId = _ownerRefUserId
	r.Set("owner_ref_user_id", _ownerRefUserId)
	return nil
}

// GetOwnerRefUserId OwnerRefUserId Getter
func (r AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) GetOwnerRefUserId() string {
	return r._ownerRefUserId
}

var poolAlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytScqyDelbillinfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytScqyDelbillinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest
func GetAlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest() *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest {
	return poolAlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest.Get().(*AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest 将 AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest(v *AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest.Put(v)
}
