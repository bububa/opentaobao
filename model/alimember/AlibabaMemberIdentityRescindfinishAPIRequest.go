package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberidentityrescindfinishAPIRequest 取消确认 API请求
// alibaba.member.identity.rescindfinish
//
// 取消确认
type AlibabamemberidentityrescindfinishAPIRequest struct {
	model.Params
	// 取消确认信息
	_rescindFinish *RescindIdentityFinishRequest
}

// NewAlibabamemberidentityrescindfinishRequest 初始化AlibabamemberidentityrescindfinishAPIRequest对象
func NewAlibabamemberidentityrescindfinishRequest() *AlibabamemberidentityrescindfinishAPIRequest {
	return &AlibabamemberidentityrescindfinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamemberidentityrescindfinishAPIRequest) GetApiMethodName() string {
	return "alibaba.member.identity.rescindfinish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamemberidentityrescindfinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamemberidentityrescindfinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRescindFinish is RescindFinish Setter
// 取消确认信息
func (r *AlibabamemberidentityrescindfinishAPIRequest) SetRescindFinish(_rescindFinish *RescindIdentityFinishRequest) error {
	r._rescindFinish = _rescindFinish
	r.Set("rescind_finish", _rescindFinish)
	return nil
}

// GetRescindFinish RescindFinish Getter
func (r AlibabamemberidentityrescindfinishAPIRequest) GetRescindFinish() *RescindIdentityFinishRequest {
	return r._rescindFinish
}
