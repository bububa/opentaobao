package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenInvoiceSearchAPIRequest 差旅申请用户搜索可用发票列表 API请求
// alitrip.btrip.open.invoice.search
//
// 差旅申请用户搜索可用发票列表
type AlitripBtripOpenInvoiceSearchAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenInvoiceRq
}

// NewAlitripBtripOpenInvoiceSearchRequest 初始化AlitripBtripOpenInvoiceSearchAPIRequest对象
func NewAlitripBtripOpenInvoiceSearchRequest() *AlitripBtripOpenInvoiceSearchAPIRequest {
	return &AlitripBtripOpenInvoiceSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenInvoiceSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenInvoiceSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.invoice.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenInvoiceSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenInvoiceSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenInvoiceSearchAPIRequest) SetRq(_rq *OpenInvoiceRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenInvoiceSearchAPIRequest) GetRq() *OpenInvoiceRq {
	return r._rq
}

var poolAlitripBtripOpenInvoiceSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenInvoiceSearchRequest()
	},
}

// GetAlitripBtripOpenInvoiceSearchRequest 从 sync.Pool 获取 AlitripBtripOpenInvoiceSearchAPIRequest
func GetAlitripBtripOpenInvoiceSearchAPIRequest() *AlitripBtripOpenInvoiceSearchAPIRequest {
	return poolAlitripBtripOpenInvoiceSearchAPIRequest.Get().(*AlitripBtripOpenInvoiceSearchAPIRequest)
}

// ReleaseAlitripBtripOpenInvoiceSearchAPIRequest 将 AlitripBtripOpenInvoiceSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenInvoiceSearchAPIRequest(v *AlitripBtripOpenInvoiceSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenInvoiceSearchAPIRequest.Put(v)
}
