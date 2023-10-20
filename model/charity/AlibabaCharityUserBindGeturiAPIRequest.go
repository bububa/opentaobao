package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUserBindGeturiAPIRequest 获取用户绑定uri API请求
// alibaba.charity.user.bind.geturi
//
// 获取用户绑定uri
type AlibabaCharityUserBindGeturiAPIRequest struct {
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

// NewAlibabaCharityUserBindGeturiRequest 初始化AlibabaCharityUserBindGeturiAPIRequest对象
func NewAlibabaCharityUserBindGeturiRequest() *AlibabaCharityUserBindGeturiAPIRequest {
	return &AlibabaCharityUserBindGeturiAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCharityUserBindGeturiAPIRequest) Reset() {
	r._features = ""
	r._platform = ""
	r._userKey = ""
	r._userNick = ""
	r._timeout = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityUserBindGeturiAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.user.bind.geturi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityUserBindGeturiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityUserBindGeturiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeatures is Features Setter
// 扩展字段
func (r *AlibabaCharityUserBindGeturiAPIRequest) SetFeatures(_features string) error {
	r._features = _features
	r.Set("features", _features)
	return nil
}

// GetFeatures Features Getter
func (r AlibabaCharityUserBindGeturiAPIRequest) GetFeatures() string {
	return r._features
}

// SetPlatform is Platform Setter
// 跳转平台的类型
func (r *AlibabaCharityUserBindGeturiAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabaCharityUserBindGeturiAPIRequest) GetPlatform() string {
	return r._platform
}

// SetUserKey is UserKey Setter
// 三方用户id
func (r *AlibabaCharityUserBindGeturiAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabaCharityUserBindGeturiAPIRequest) GetUserKey() string {
	return r._userKey
}

// SetUserNick is UserNick Setter
// 三方用户昵称
func (r *AlibabaCharityUserBindGeturiAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r AlibabaCharityUserBindGeturiAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetTimeout is Timeout Setter
// 链接超时时间
func (r *AlibabaCharityUserBindGeturiAPIRequest) SetTimeout(_timeout int64) error {
	r._timeout = _timeout
	r.Set("timeout", _timeout)
	return nil
}

// GetTimeout Timeout Getter
func (r AlibabaCharityUserBindGeturiAPIRequest) GetTimeout() int64 {
	return r._timeout
}

var poolAlibabaCharityUserBindGeturiAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCharityUserBindGeturiRequest()
	},
}

// GetAlibabaCharityUserBindGeturiRequest 从 sync.Pool 获取 AlibabaCharityUserBindGeturiAPIRequest
func GetAlibabaCharityUserBindGeturiAPIRequest() *AlibabaCharityUserBindGeturiAPIRequest {
	return poolAlibabaCharityUserBindGeturiAPIRequest.Get().(*AlibabaCharityUserBindGeturiAPIRequest)
}

// ReleaseAlibabaCharityUserBindGeturiAPIRequest 将 AlibabaCharityUserBindGeturiAPIRequest 放入 sync.Pool
func ReleaseAlibabaCharityUserBindGeturiAPIRequest(v *AlibabaCharityUserBindGeturiAPIRequest) {
	v.Reset()
	poolAlibabaCharityUserBindGeturiAPIRequest.Put(v)
}
