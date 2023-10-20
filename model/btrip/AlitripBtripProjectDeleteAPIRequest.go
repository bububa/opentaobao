package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripprojectdeleteAPIRequest 删除项目 API请求
// alitrip.btrip.project.delete
//
// 删除项目
type AlitripbtripprojectdeleteAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripbtripprojectdeleteRequest 初始化AlitripbtripprojectdeleteAPIRequest对象
func NewAlitripbtripprojectdeleteRequest() *AlitripbtripprojectdeleteAPIRequest {
	return &AlitripbtripprojectdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripprojectdeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripprojectdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripprojectdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripprojectdeleteAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripprojectdeleteAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}
