package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrbillcheckAPIRequest 疫苗追溯验证 API请求
// alibaba.alihealth.drug.kyt.dr.billcheck
//
// 各级疾控在入库完成后，需要做追溯信息验证
type AlibabaalihealthdrugkytdrbillcheckAPIRequest struct {
	model.Params
	// 调用企业ID
	_refEntId string
	// 单据编号
	_billCode string
	// 单据类型
	_billType string
	// 单据企业refEntId
	_owerRefEntId string
}

// NewAlibabaalihealthdrugkytdrbillcheckRequest 初始化AlibabaalihealthdrugkytdrbillcheckAPIRequest对象
func NewAlibabaalihealthdrugkytdrbillcheckRequest() *AlibabaalihealthdrugkytdrbillcheckAPIRequest {
	return &AlibabaalihealthdrugkytdrbillcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrbillcheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.billcheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrbillcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrbillcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用企业ID
func (r *AlibabaalihealthdrugkytdrbillcheckAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrbillcheckAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaalihealthdrugkytdrbillcheckAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytdrbillcheckAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaalihealthdrugkytdrbillcheckAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytdrbillcheckAPIRequest) GetBillType() string {
	return r._billType
}

// SetOwerRefEntId is OwerRefEntId Setter
// 单据企业refEntId
func (r *AlibabaalihealthdrugkytdrbillcheckAPIRequest) SetOwerRefEntId(_owerRefEntId string) error {
	r._owerRefEntId = _owerRefEntId
	r.Set("ower_ref_ent_id", _owerRefEntId)
	return nil
}

// GetOwerRefEntId OwerRefEntId Getter
func (r AlibabaalihealthdrugkytdrbillcheckAPIRequest) GetOwerRefEntId() string {
	return r._owerRefEntId
}
