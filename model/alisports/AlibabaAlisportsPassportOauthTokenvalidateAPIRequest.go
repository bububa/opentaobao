package alisports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportOauthTokenvalidateAPIRequest 阿里体育会员系统--获取登录信息接口 API请求
// alibaba.alisports.passport.oauth.tokenvalidate
//
// 阿里体育会员系统--获取登录信息接口
type AlibabaAlisportsPassportOauthTokenvalidateAPIRequest struct {
	model.Params
	// 登录时返回给前端的token
	_token string
	// 时间戳
	_alispTime string
	// 应用的appkey
	_alispAppKey string
	// 参数加密之后的串
	_alispSign string
}

// NewAlibabaAlisportsPassportOauthTokenvalidateRequest 初始化AlibabaAlisportsPassportOauthTokenvalidateAPIRequest对象
func NewAlibabaAlisportsPassportOauthTokenvalidateRequest() *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest {
	return &AlibabaAlisportsPassportOauthTokenvalidateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) Reset() {
	r._token = ""
	r._alispTime = ""
	r._alispAppKey = ""
	r._alispSign = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.oauth.tokenvalidate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 登录时返回给前端的token
func (r *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) GetToken() string {
	return r._token
}

// SetAlispTime is AlispTime Setter
// 时间戳
func (r *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispAppKey is AlispAppKey Setter
// 应用的appkey
func (r *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispSign is AlispSign Setter
// 参数加密之后的串
func (r *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) GetAlispSign() string {
	return r._alispSign
}

var poolAlibabaAlisportsPassportOauthTokenvalidateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlisportsPassportOauthTokenvalidateRequest()
	},
}

// GetAlibabaAlisportsPassportOauthTokenvalidateRequest 从 sync.Pool 获取 AlibabaAlisportsPassportOauthTokenvalidateAPIRequest
func GetAlibabaAlisportsPassportOauthTokenvalidateAPIRequest() *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest {
	return poolAlibabaAlisportsPassportOauthTokenvalidateAPIRequest.Get().(*AlibabaAlisportsPassportOauthTokenvalidateAPIRequest)
}

// ReleaseAlibabaAlisportsPassportOauthTokenvalidateAPIRequest 将 AlibabaAlisportsPassportOauthTokenvalidateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlisportsPassportOauthTokenvalidateAPIRequest(v *AlibabaAlisportsPassportOauthTokenvalidateAPIRequest) {
	v.Reset()
	poolAlibabaAlisportsPassportOauthTokenvalidateAPIRequest.Put(v)
}
