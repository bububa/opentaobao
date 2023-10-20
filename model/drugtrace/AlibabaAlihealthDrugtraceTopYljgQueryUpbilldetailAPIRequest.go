package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 API请求
// alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail
//
// 根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 本企业refEntId
	_refEntId string
}

// NewAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest 初始化AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest() *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 本企业refEntId
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}
