package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripProjectAddAPIRequest 添加项目 API请求
// alitrip.btrip.project.add
//
// 添加项目
type AlitripBtripProjectAddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripBtripProjectAddRequest 初始化AlitripBtripProjectAddAPIRequest对象
func NewAlitripBtripProjectAddRequest() *AlitripBtripProjectAddAPIRequest {
	return &AlitripBtripProjectAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripProjectAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripProjectAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripProjectAddAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripProjectAddAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}
