package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripprojectmodifyAPIRequest 变更项目 API请求
// alitrip.btrip.project.modify
//
// 变更项目
type AlitripbtripprojectmodifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripbtripprojectmodifyRequest 初始化AlitripbtripprojectmodifyAPIRequest对象
func NewAlitripbtripprojectmodifyRequest() *AlitripbtripprojectmodifyAPIRequest {
	return &AlitripbtripprojectmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripprojectmodifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripprojectmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripprojectmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripprojectmodifyAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripprojectmodifyAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}
