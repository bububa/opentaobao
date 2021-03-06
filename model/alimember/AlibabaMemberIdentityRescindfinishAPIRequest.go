package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIdentityRescindfinishAPIRequest 取消确认 API请求
// alibaba.member.identity.rescindfinish
//
// 取消确认
type AlibabaMemberIdentityRescindfinishAPIRequest struct {
	model.Params
	// 取消确认信息
	_rescindFinish *RescindIdentityFinishRequest
}

// NewAlibabaMemberIdentityRescindfinishRequest 初始化AlibabaMemberIdentityRescindfinishAPIRequest对象
func NewAlibabaMemberIdentityRescindfinishRequest() *AlibabaMemberIdentityRescindfinishAPIRequest {
	return &AlibabaMemberIdentityRescindfinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberIdentityRescindfinishAPIRequest) GetApiMethodName() string {
	return "alibaba.member.identity.rescindfinish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberIdentityRescindfinishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRescindFinish is RescindFinish Setter
// 取消确认信息
func (r *AlibabaMemberIdentityRescindfinishAPIRequest) SetRescindFinish(_rescindFinish *RescindIdentityFinishRequest) error {
	r._rescindFinish = _rescindFinish
	r.Set("rescind_finish", _rescindFinish)
	return nil
}

// GetRescindFinish RescindFinish Getter
func (r AlibabaMemberIdentityRescindfinishAPIRequest) GetRescindFinish() *RescindIdentityFinishRequest {
	return r._rescindFinish
}
