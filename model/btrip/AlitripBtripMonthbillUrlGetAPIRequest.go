package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripMonthbillUrlGetAPIRequest 月账单数据查询 API请求
// alitrip.btrip.monthbill.url.get
//
// 月账单数据查询
type AlitripBtripMonthbillUrlGetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenAccountRq
}

// NewAlitripBtripMonthbillUrlGetRequest 初始化AlitripBtripMonthbillUrlGetAPIRequest对象
func NewAlitripBtripMonthbillUrlGetRequest() *AlitripBtripMonthbillUrlGetAPIRequest {
	return &AlitripBtripMonthbillUrlGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripMonthbillUrlGetAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripMonthbillUrlGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.monthbill.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripMonthbillUrlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripMonthbillUrlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripMonthbillUrlGetAPIRequest) SetRq(_rq *OpenAccountRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripMonthbillUrlGetAPIRequest) GetRq() *OpenAccountRq {
	return r._rq
}

var poolAlitripBtripMonthbillUrlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripMonthbillUrlGetRequest()
	},
}

// GetAlitripBtripMonthbillUrlGetRequest 从 sync.Pool 获取 AlitripBtripMonthbillUrlGetAPIRequest
func GetAlitripBtripMonthbillUrlGetAPIRequest() *AlitripBtripMonthbillUrlGetAPIRequest {
	return poolAlitripBtripMonthbillUrlGetAPIRequest.Get().(*AlitripBtripMonthbillUrlGetAPIRequest)
}

// ReleaseAlitripBtripMonthbillUrlGetAPIRequest 将 AlitripBtripMonthbillUrlGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripMonthbillUrlGetAPIRequest(v *AlitripBtripMonthbillUrlGetAPIRequest) {
	v.Reset()
	poolAlitripBtripMonthbillUrlGetAPIRequest.Put(v)
}
