package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopCommonapplyGetAPIRequest 商旅审批单通用查询接口 API请求
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
type AlitripBtripCorpopCommonapplyGetAPIRequest struct {
	model.Params
	// 请求入参
	_rq *OpenIsvSearchRQ
}

// NewAlitripBtripCorpopCommonapplyGetRequest 初始化AlitripBtripCorpopCommonapplyGetAPIRequest对象
func NewAlitripBtripCorpopCommonapplyGetRequest() *AlitripBtripCorpopCommonapplyGetAPIRequest {
	return &AlitripBtripCorpopCommonapplyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopCommonapplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.commonapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopCommonapplyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 请求入参
func (r *AlitripBtripCorpopCommonapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRQ) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopCommonapplyGetAPIRequest) GetRq() *OpenIsvSearchRQ {
	return r._rq
}
