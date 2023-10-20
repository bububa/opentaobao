package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychaintraincityAPIRequest 火车站数据查询 API请求
// alitrip.btrip.supplychain.train.city
//
// 火车站数据查询
type AlitripbtripsupplychaintraincityAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenSuggestRq
}

// NewAlitripbtripsupplychaintraincityRequest 初始化AlitripbtripsupplychaintraincityAPIRequest对象
func NewAlitripbtripsupplychaintraincityRequest() *AlitripbtripsupplychaintraincityAPIRequest {
	return &AlitripbtripsupplychaintraincityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychaintraincityAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.city"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychaintraincityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychaintraincityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripsupplychaintraincityAPIRequest) SetRq(_rq *OpenSuggestRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychaintraincityAPIRequest) GetRq() *OpenSuggestRq {
	return r._rq
}
