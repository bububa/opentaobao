package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest 同步供应商校准后的配额-二级物料 API请求
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload
//
// 同步供应商校准后的配额-二级物料
type AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest struct {
	model.Params
	// 对象
	_currentQuoteRawRequest *AbstractRequest
}

// NewAlibabatmallgeniescpplancorrectsupplierquoterawuploadRequest 初始化AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest对象
func NewAlibabatmallgeniescpplancorrectsupplierquoterawuploadRequest() *AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest {
	return &AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCurrentQuoteRawRequest is CurrentQuoteRawRequest Setter
// 对象
func (r *AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest) SetCurrentQuoteRawRequest(_currentQuoteRawRequest *AbstractRequest) error {
	r._currentQuoteRawRequest = _currentQuoteRawRequest
	r.Set("current_quote_raw_request", _currentQuoteRawRequest)
	return nil
}

// GetCurrentQuoteRawRequest CurrentQuoteRawRequest Getter
func (r AlibabatmallgeniescpplancorrectsupplierquoterawuploadAPIRequest) GetCurrentQuoteRawRequest() *AbstractRequest {
	return r._currentQuoteRawRequest
}
