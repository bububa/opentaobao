package alimember

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIdentitySignfinishAPIRequest 签约确认 API请求
// alibaba.member.identity.signfinish
//
// 签约确认
type AlibabaMemberIdentitySignfinishAPIRequest struct {
	model.Params
	// 签约确认信息
	_signFinish *SignIdentityFinishRequest
}

// NewAlibabaMemberIdentitySignfinishRequest 初始化AlibabaMemberIdentitySignfinishAPIRequest对象
func NewAlibabaMemberIdentitySignfinishRequest() *AlibabaMemberIdentitySignfinishAPIRequest {
	return &AlibabaMemberIdentitySignfinishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberIdentitySignfinishAPIRequest) Reset() {
	r._signFinish = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberIdentitySignfinishAPIRequest) GetApiMethodName() string {
	return "alibaba.member.identity.signfinish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberIdentitySignfinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberIdentitySignfinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSignFinish is SignFinish Setter
// 签约确认信息
func (r *AlibabaMemberIdentitySignfinishAPIRequest) SetSignFinish(_signFinish *SignIdentityFinishRequest) error {
	r._signFinish = _signFinish
	r.Set("sign_finish", _signFinish)
	return nil
}

// GetSignFinish SignFinish Getter
func (r AlibabaMemberIdentitySignfinishAPIRequest) GetSignFinish() *SignIdentityFinishRequest {
	return r._signFinish
}

var poolAlibabaMemberIdentitySignfinishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberIdentitySignfinishRequest()
	},
}

// GetAlibabaMemberIdentitySignfinishRequest 从 sync.Pool 获取 AlibabaMemberIdentitySignfinishAPIRequest
func GetAlibabaMemberIdentitySignfinishAPIRequest() *AlibabaMemberIdentitySignfinishAPIRequest {
	return poolAlibabaMemberIdentitySignfinishAPIRequest.Get().(*AlibabaMemberIdentitySignfinishAPIRequest)
}

// ReleaseAlibabaMemberIdentitySignfinishAPIRequest 将 AlibabaMemberIdentitySignfinishAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberIdentitySignfinishAPIRequest(v *AlibabaMemberIdentitySignfinishAPIRequest) {
	v.Reset()
	poolAlibabaMemberIdentitySignfinishAPIRequest.Put(v)
}
