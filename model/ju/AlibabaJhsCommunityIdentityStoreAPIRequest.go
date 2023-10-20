package ju

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityIdentityStoreAPIRequest 用户信息存储 API请求
// alibaba.jhs.community.identity.store
//
// 用户信息存储
type AlibabaJhsCommunityIdentityStoreAPIRequest struct {
	model.Params
	// 会话token
	_token string
	// 用户昵称
	_nickName string
	// 用户头像
	_avatarUrl string
}

// NewAlibabaJhsCommunityIdentityStoreRequest 初始化AlibabaJhsCommunityIdentityStoreAPIRequest对象
func NewAlibabaJhsCommunityIdentityStoreRequest() *AlibabaJhsCommunityIdentityStoreAPIRequest {
	return &AlibabaJhsCommunityIdentityStoreAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJhsCommunityIdentityStoreAPIRequest) Reset() {
	r._token = ""
	r._nickName = ""
	r._avatarUrl = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunityIdentityStoreAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.identity.store"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunityIdentityStoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunityIdentityStoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 会话token
func (r *AlibabaJhsCommunityIdentityStoreAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaJhsCommunityIdentityStoreAPIRequest) GetToken() string {
	return r._token
}

// SetNickName is NickName Setter
// 用户昵称
func (r *AlibabaJhsCommunityIdentityStoreAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// GetNickName NickName Getter
func (r AlibabaJhsCommunityIdentityStoreAPIRequest) GetNickName() string {
	return r._nickName
}

// SetAvatarUrl is AvatarUrl Setter
// 用户头像
func (r *AlibabaJhsCommunityIdentityStoreAPIRequest) SetAvatarUrl(_avatarUrl string) error {
	r._avatarUrl = _avatarUrl
	r.Set("avatar_url", _avatarUrl)
	return nil
}

// GetAvatarUrl AvatarUrl Getter
func (r AlibabaJhsCommunityIdentityStoreAPIRequest) GetAvatarUrl() string {
	return r._avatarUrl
}

var poolAlibabaJhsCommunityIdentityStoreAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJhsCommunityIdentityStoreRequest()
	},
}

// GetAlibabaJhsCommunityIdentityStoreRequest 从 sync.Pool 获取 AlibabaJhsCommunityIdentityStoreAPIRequest
func GetAlibabaJhsCommunityIdentityStoreAPIRequest() *AlibabaJhsCommunityIdentityStoreAPIRequest {
	return poolAlibabaJhsCommunityIdentityStoreAPIRequest.Get().(*AlibabaJhsCommunityIdentityStoreAPIRequest)
}

// ReleaseAlibabaJhsCommunityIdentityStoreAPIRequest 将 AlibabaJhsCommunityIdentityStoreAPIRequest 放入 sync.Pool
func ReleaseAlibabaJhsCommunityIdentityStoreAPIRequest(v *AlibabaJhsCommunityIdentityStoreAPIRequest) {
	v.Reset()
	poolAlibabaJhsCommunityIdentityStoreAPIRequest.Put(v)
}
