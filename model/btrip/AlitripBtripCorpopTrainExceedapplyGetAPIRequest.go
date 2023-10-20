package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopTrainExceedapplyGetAPIRequest 商旅火车票第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
type AlitripBtripCorpopTrainExceedapplyGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRq
}

// NewAlitripBtripCorpopTrainExceedapplyGetRequest 初始化AlitripBtripCorpopTrainExceedapplyGetAPIRequest对象
func NewAlitripBtripCorpopTrainExceedapplyGetRequest() *AlitripBtripCorpopTrainExceedapplyGetAPIRequest {
	return &AlitripBtripCorpopTrainExceedapplyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.train.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopTrainExceedapplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopTrainExceedapplyGetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
