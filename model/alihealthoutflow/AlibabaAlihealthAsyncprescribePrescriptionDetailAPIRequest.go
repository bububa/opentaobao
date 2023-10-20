package alihealthoutflow

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) Reset() {
	r._detailRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.asyncprescribe.prescription.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailRequest is DetailRequest Setter
// 入参
func (r *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) SetDetailRequest(_detailRequest *AsyncPrescribeDetailRequest) error {
	r._detailRequest = _detailRequest
	r.Set("detail_request", _detailRequest)
	return nil
}

// GetDetailRequest DetailRequest Getter
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) GetDetailRequest() *AsyncPrescribeDetailRequest {
	return r._detailRequest
}

var poolAlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthAsyncprescribePrescriptionDetailRequest()
	},
}

// GetAlibabaAlihealthAsyncprescribePrescriptionDetailRequest 从 sync.Pool 获取 AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest
func GetAlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest() *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest {
	return poolAlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest.Get().(*AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest)
}

// ReleaseAlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest 将 AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest(v *AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest.Put(v)
}
