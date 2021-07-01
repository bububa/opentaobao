package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripProjectDeleteAPIRequest
删除项目 API请求
alitrip.btrip.project.delete

删除项目 */
type AlitripBtripProjectDeleteAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// NewAlitripBtripProjectDeleteRequest 初始化AlitripBtripProjectDeleteAPIRequest对象
func NewAlitripBtripProjectDeleteRequest() *AlitripBtripProjectDeleteAPIRequest {
	return &AlitripBtripProjectDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripProjectDeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.project.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripProjectDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripProjectDeleteAPIRequest) SetRq(_rq *OpenProjectRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripProjectDeleteAPIRequest) GetRq() *OpenProjectRq {
	return r._rq
}
