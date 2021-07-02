package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRxPrescriptionDetailAPIRequest 处方详情 API请求
// alibaba.alihealth.rx.prescription.detail
//
// 获取处方结构化信息
type AlibabaAlihealthRxPrescriptionDetailAPIRequest struct {
	model.Params
	// 查询参数
	_query *RxPrescriptionQuery
}

// NewAlibabaAlihealthRxPrescriptionDetailRequest 初始化AlibabaAlihealthRxPrescriptionDetailAPIRequest对象
func NewAlibabaAlihealthRxPrescriptionDetailRequest() *AlibabaAlihealthRxPrescriptionDetailAPIRequest {
	return &AlibabaAlihealthRxPrescriptionDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRxPrescriptionDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.rx.prescription.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRxPrescriptionDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuery is Query Setter
// 查询参数
func (r *AlibabaAlihealthRxPrescriptionDetailAPIRequest) SetQuery(_query *RxPrescriptionQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaAlihealthRxPrescriptionDetailAPIRequest) GetQuery() *RxPrescriptionQuery {
	return r._query
}
