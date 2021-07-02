package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest 异步开方处方详情 API请求
// alibaba.alihealth.asyncprescribe.prescription.detail
//
// 异步开方处方查询
type AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest struct {
	model.Params
	// 入参
	_detailRequest *AsyncPrescribeDetailRequest
}

// NewAlibabaAlihealthAsyncprescribePrescriptionDetailRequest 初始化AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest对象
func NewAlibabaAlihealthAsyncprescribePrescriptionDetailRequest() *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest {
	return &AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.asyncprescribe.prescription.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DetailRequest Setter
// 入参
func (r *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) SetDetailRequest(_detailRequest *AsyncPrescribeDetailRequest) error {
	r._detailRequest = _detailRequest
	r.Set("detail_request", _detailRequest)
	return nil
}

// Get DetailRequest Getter
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) GetDetailRequest() *AsyncPrescribeDetailRequest {
	return r._detailRequest
}
