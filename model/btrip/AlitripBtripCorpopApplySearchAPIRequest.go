package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplySearchAPIRequest 【商旅】搜索审批单列表 API请求
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
type AlitripBtripCorpopApplySearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenIsvSearchRQ
}

// NewAlitripBtripCorpopApplySearchRequest 初始化AlitripBtripCorpopApplySearchAPIRequest对象
func NewAlitripBtripCorpopApplySearchRequest() *AlitripBtripCorpopApplySearchAPIRequest {
	return &AlitripBtripCorpopApplySearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplySearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplySearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplySearchAPIRequest) SetRq(_rq *OpenIsvSearchRQ) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopApplySearchAPIRequest) GetRq() *OpenIsvSearchRQ {
	return r._rq
}
