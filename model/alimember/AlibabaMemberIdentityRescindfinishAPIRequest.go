package alimember

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberIdentityRescindfinishAPIRequest) Reset() {
	r._rescindFinish = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberIdentityRescindfinishAPIRequest) GetApiMethodName() string {
	return "alibaba.member.identity.rescindfinish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberIdentityRescindfinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberIdentityRescindfinishAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaMemberIdentityRescindfinishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberIdentityRescindfinishRequest()
	},
}

// GetAlibabaMemberIdentityRescindfinishRequest 从 sync.Pool 获取 AlibabaMemberIdentityRescindfinishAPIRequest
func GetAlibabaMemberIdentityRescindfinishAPIRequest() *AlibabaMemberIdentityRescindfinishAPIRequest {
	return poolAlibabaMemberIdentityRescindfinishAPIRequest.Get().(*AlibabaMemberIdentityRescindfinishAPIRequest)
}

// ReleaseAlibabaMemberIdentityRescindfinishAPIRequest 将 AlibabaMemberIdentityRescindfinishAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberIdentityRescindfinishAPIRequest(v *AlibabaMemberIdentityRescindfinishAPIRequest) {
	v.Reset()
	poolAlibabaMemberIdentityRescindfinishAPIRequest.Put(v)
}
