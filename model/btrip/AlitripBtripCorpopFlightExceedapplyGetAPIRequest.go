package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopFlightExceedapplyGetAPIRequest 商旅机票第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
type AlitripBtripCorpopFlightExceedapplyGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRQ
}

// NewAlitripBtripCorpopFlightExceedapplyGetRequest 初始化AlitripBtripCorpopFlightExceedapplyGetAPIRequest对象
func NewAlitripBtripCorpopFlightExceedapplyGetRequest() *AlitripBtripCorpopFlightExceedapplyGetAPIRequest {
	return &AlitripBtripCorpopFlightExceedapplyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopFlightExceedapplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.flight.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopFlightExceedapplyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopFlightExceedapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRQ) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopFlightExceedapplyGetAPIRequest) GetRq() *OpenIsvSearchRQ {
	return r._rq
}
