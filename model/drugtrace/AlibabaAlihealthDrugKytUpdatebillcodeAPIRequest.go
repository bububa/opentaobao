package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytupdatebillcodeAPIRequest 零售修改出入库单追溯码 API请求
// alibaba.alihealth.drug.kyt.updatebillcode
//
// 零售修改出入库单追溯码
type AlibabaalihealthdrugkytupdatebillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// 企业ID
	_refEntId string
	// 操作人ID
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
}

// NewAlibabaalihealthdrugkytupdatebillcodeRequest 初始化AlibabaalihealthdrugkytupdatebillcodeAPIRequest对象
func NewAlibabaalihealthdrugkytupdatebillcodeRequest() *AlibabaalihealthdrugkytupdatebillcodeAPIRequest {
	return &AlibabaalihealthdrugkytupdatebillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.updatebillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 追溯码
func (r *AlibabaalihealthdrugkytupdatebillcodeAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytupdatebillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetIcCode is IcCode Setter
// 操作人ID
func (r *AlibabaalihealthdrugkytupdatebillcodeAPIRequest) SetIcCode(_icCode string) error {
	r._icCode = _icCode
	r.Set("ic_code", _icCode)
	return nil
}

// GetIcCode IcCode Getter
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetIcCode() string {
	return r._icCode
}

// SetBillId is BillId Setter
// 单据ID
func (r *AlibabaalihealthdrugkytupdatebillcodeAPIRequest) SetBillId(_billId string) error {
	r._billId = _billId
	r.Set("bill_id", _billId)
	return nil
}

// GetBillId BillId Getter
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetBillId() string {
	return r._billId
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaalihealthdrugkytupdatebillcodeAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytupdatebillcodeAPIRequest) GetBillType() string {
	return r._billType
}
