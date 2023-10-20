package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest 同步供应商校准后的配额-二级物料 API请求
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload
//
// 同步供应商校准后的配额-二级物料
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest struct {
	model.Params
	// 对象
	_currentQuoteRawRequest *AbstractRequest
}

// NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest 初始化AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) Reset() {
	r._currentQuoteRawRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCurrentQuoteRawRequest is CurrentQuoteRawRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) SetCurrentQuoteRawRequest(_currentQuoteRawRequest *AbstractRequest) error {
	r._currentQuoteRawRequest = _currentQuoteRawRequest
	r.Set("current_quote_raw_request", _currentQuoteRawRequest)
	return nil
}

// GetCurrentQuoteRawRequest CurrentQuoteRawRequest Getter
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) GetCurrentQuoteRawRequest() *AbstractRequest {
	return r._currentQuoteRawRequest
}

var poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest
func GetAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest 将 AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest(v *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest.Put(v)
}
