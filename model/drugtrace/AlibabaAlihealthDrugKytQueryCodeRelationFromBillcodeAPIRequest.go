package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest 根据单据号码查询码单据详情和码信息 API请求
// alibaba.alihealth.drug.kyt.query.code.relation.from.billcode
//
// 根据单据号码查询码单据详情和码信息
type AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 企业refEntId
	_refEntId string
}

// NewAlibabaalihealthdrugkytquerycoderelationfrombillcodeRequest 初始化AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest对象
func NewAlibabaalihealthdrugkytquerycoderelationfrombillcodeRequest() *AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest {
	return &AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.code.relation.from.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
