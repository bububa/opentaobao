package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyAddAPIRequest 【商旅】isv添加审批单 API请求
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
type AlitripBtripCorpopApplyAddAPIRequest struct {
	model.Params
	// 请求参数
	_rq *OpenApiApplyRq
}

// NewAlitripBtripCorpopApplyAddRequest 初始化AlitripBtripCorpopApplyAddAPIRequest对象
func NewAlitripBtripCorpopApplyAddRequest() *AlitripBtripCorpopApplyAddAPIRequest {
	return &AlitripBtripCorpopApplyAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求参数
func (r *AlitripBtripCorpopApplyAddAPIRequest) SetRq(_rq *OpenApiApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCorpopApplyAddAPIRequest) GetRq() *OpenApiApplyRq {
	return r._rq
}
