package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
