package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest 异步开方处方详情 API请求
// alibaba.alihealth.asyncprescribe.prescription.detail
//
// 异步开方处方查询
type AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest struct {
	model.Params
	// 入参
	_detailRequest *AsyncPrescribeDetailRequest
}

// NewAlibabaalihealthasyncprescribeprescriptiondetailRequest 初始化AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest对象
func NewAlibabaalihealthasyncprescribeprescriptiondetailRequest() *AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest {
	return &AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.asyncprescribe.prescription.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailRequest is DetailRequest Setter
// 入参
func (r *AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest) SetDetailRequest(_detailRequest *AsyncPrescribeDetailRequest) error {
	r._detailRequest = _detailRequest
	r.Set("detail_request", _detailRequest)
	return nil
}

// GetDetailRequest DetailRequest Getter
func (r AlibabaalihealthasyncprescribeprescriptiondetailAPIRequest) GetDetailRequest() *AsyncPrescribeDetailRequest {
	return r._detailRequest
}
