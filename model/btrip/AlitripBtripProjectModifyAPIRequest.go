package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripProjectModifyAPIRequest 变更项目 API请求
// alitrip.btrip.project.modify
//
// 变更项目
type AlitripBtripProjectModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripBtripProjectModifyRequest 初始化AlitripBtripProjectModifyAPIRequest对象
func NewAlitripBtripProjectModifyRequest() *AlitripBtripProjectModifyAPIRequest {
	return &AlitripBtripProjectModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripProjectModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripProjectModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripProjectModifyAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripProjectModifyAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}
