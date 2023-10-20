package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdestbillcheckAPIRequest 直调审批 API请求
// alibaba.alihealth.drug.kyt.destbill.check
//
// 为药企提供直调单据审批操作
type AlibabaalihealthdrugkytdestbillcheckAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 单据号
	_billCode string
	// 审核状态，'Y'审批通过 'N' 审批不通过
	_checkType string
}

// NewAlibabaalihealthdrugkytdestbillcheckRequest 初始化AlibabaalihealthdrugkytdestbillcheckAPIRequest对象
func NewAlibabaalihealthdrugkytdestbillcheckRequest() *AlibabaalihealthdrugkytdestbillcheckAPIRequest {
	return &AlibabaalihealthdrugkytdestbillcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdestbillcheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.destbill.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdestbillcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdestbillcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytdestbillcheckAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdestbillcheckAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaalihealthdrugkytdestbillcheckAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytdestbillcheckAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetCheckType is CheckType Setter
// 审核状态，&#39;Y&#39;审批通过 &#39;N&#39; 审批不通过
func (r *AlibabaalihealthdrugkytdestbillcheckAPIRequest) SetCheckType(_checkType string) error {
	r._checkType = _checkType
	r.Set("check_type", _checkType)
	return nil
}

// GetCheckType CheckType Getter
func (r AlibabaalihealthdrugkytdestbillcheckAPIRequest) GetCheckType() string {
	return r._checkType
}
