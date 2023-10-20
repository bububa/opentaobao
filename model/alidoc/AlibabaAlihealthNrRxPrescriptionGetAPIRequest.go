package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrRxPrescriptionGetAPIRequest 搜索处方详情 API请求
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
type AlibabaAlihealthNrRxPrescriptionGetAPIRequest struct {
	model.Params
	// 查询参数
	_query *PrescriptionQueryDto
}

// NewAlibabaAlihealthNrRxPrescriptionGetRequest 初始化AlibabaAlihealthNrRxPrescriptionGetAPIRequest对象
func NewAlibabaAlihealthNrRxPrescriptionGetRequest() *AlibabaAlihealthNrRxPrescriptionGetAPIRequest {
	return &AlibabaAlihealthNrRxPrescriptionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrRxPrescriptionGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.rx.prescription.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrRxPrescriptionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthNrRxPrescriptionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *AlibabaAlihealthNrRxPrescriptionGetAPIRequest) SetQuery(_query *PrescriptionQueryDto) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaAlihealthNrRxPrescriptionGetAPIRequest) GetQuery() *PrescriptionQueryDto {
	return r._query
}
