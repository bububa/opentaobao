package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharityuserbindgeturiAPIRequest 获取用户绑定uri API请求
// alibaba.charity.user.bind.geturi
//
// 获取用户绑定uri
type AlibabacharityuserbindgeturiAPIRequest struct {
	model.Params
	// 扩展字段
	_features string
	// 跳转平台的类型
	_platform string
	// 三方用户id
	_userKey string
	// 三方用户昵称
	_userNick string
	// 链接超时时间
	_timeout int64
}

// NewAlibabacharityuserbindgeturiRequest 初始化AlibabacharityuserbindgeturiAPIRequest对象
func NewAlibabacharityuserbindgeturiRequest() *AlibabacharityuserbindgeturiAPIRequest {
	return &AlibabacharityuserbindgeturiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharityuserbindgeturiAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.user.bind.geturi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharityuserbindgeturiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharityuserbindgeturiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeatures is Features Setter
// 扩展字段
func (r *AlibabacharityuserbindgeturiAPIRequest) SetFeatures(_features string) error {
	r._features = _features
	r.Set("features", _features)
	return nil
}

// GetFeatures Features Getter
func (r AlibabacharityuserbindgeturiAPIRequest) GetFeatures() string {
	return r._features
}

// SetPlatform is Platform Setter
// 跳转平台的类型
func (r *AlibabacharityuserbindgeturiAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabacharityuserbindgeturiAPIRequest) GetPlatform() string {
	return r._platform
}

// SetUserKey is UserKey Setter
// 三方用户id
func (r *AlibabacharityuserbindgeturiAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabacharityuserbindgeturiAPIRequest) GetUserKey() string {
	return r._userKey
}

// SetUserNick is UserNick Setter
// 三方用户昵称
func (r *AlibabacharityuserbindgeturiAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r AlibabacharityuserbindgeturiAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetTimeout is Timeout Setter
// 链接超时时间
func (r *AlibabacharityuserbindgeturiAPIRequest) SetTimeout(_timeout int64) error {
	r._timeout = _timeout
	r.Set("timeout", _timeout)
	return nil
}

// GetTimeout Timeout Getter
func (r AlibabacharityuserbindgeturiAPIRequest) GetTimeout() int64 {
	return r._timeout
}
