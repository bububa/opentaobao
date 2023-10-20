package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest 根据单据编号查询单据明细 API请求
// alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode
//
// 根据单据编号查询单据明细
type AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest struct {
	model.Params
	// 单据号
	_billCode string
	// 企业id
	_refEntId string
}

// NewAlibabaalihealthdrugkytqueryspeciavaccinbillcodeRequest 初始化AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest对象
func NewAlibabaalihealthdrugkytqueryspeciavaccinbillcodeRequest() *AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest {
	return &AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
