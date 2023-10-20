package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripprojectaddAPIRequest 添加项目 API请求
// alitrip.btrip.project.add
//
// 添加项目
type AlitripbtripprojectaddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripbtripprojectaddRequest 初始化AlitripbtripprojectaddAPIRequest对象
func NewAlitripbtripprojectaddRequest() *AlitripbtripprojectaddAPIRequest {
	return &AlitripbtripprojectaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripprojectaddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripprojectaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripprojectaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripprojectaddAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripprojectaddAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}
