package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripmonthbillurlgetAPIRequest 月账单数据查询 API请求
// alitrip.btrip.monthbill.url.get
//
// 月账单数据查询
type AlitripbtripmonthbillurlgetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenAccountRq
}

// NewAlitripbtripmonthbillurlgetRequest 初始化AlitripbtripmonthbillurlgetAPIRequest对象
func NewAlitripbtripmonthbillurlgetRequest() *AlitripbtripmonthbillurlgetAPIRequest {
	return &AlitripbtripmonthbillurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripmonthbillurlgetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.monthbill.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripmonthbillurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripmonthbillurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripmonthbillurlgetAPIRequest) SetRq(_rq *OpenAccountRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripmonthbillurlgetAPIRequest) GetRq() *OpenAccountRq {
	return r._rq
}
