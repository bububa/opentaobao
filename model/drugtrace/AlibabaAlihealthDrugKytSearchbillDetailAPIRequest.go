package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytsearchbilldetailAPIRequest 查询单据详情 API请求
// alibaba.alihealth.drug.kyt.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
type AlibabaalihealthdrugkytsearchbilldetailAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 企业refEntId
	_refEntId string
}

// NewAlibabaalihealthdrugkytsearchbilldetailRequest 初始化AlibabaalihealthdrugkytsearchbilldetailAPIRequest对象
func NewAlibabaalihealthdrugkytsearchbilldetailRequest() *AlibabaalihealthdrugkytsearchbilldetailAPIRequest {
	return &AlibabaalihealthdrugkytsearchbilldetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytsearchbilldetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.searchbill.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytsearchbilldetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytsearchbilldetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaalihealthdrugkytsearchbilldetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytsearchbilldetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaalihealthdrugkytsearchbilldetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytsearchbilldetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}
