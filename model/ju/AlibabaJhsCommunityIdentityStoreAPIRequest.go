package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunityidentitystoreAPIRequest 用户信息存储 API请求
// alibaba.jhs.community.identity.store
//
// 用户信息存储
type AlibabajhscommunityidentitystoreAPIRequest struct {
	model.Params
	// 会话token
	_token string
	// 用户昵称
	_nickName string
	// 用户头像
	_avatarUrl string
}

// NewAlibabajhscommunityidentitystoreRequest 初始化AlibabajhscommunityidentitystoreAPIRequest对象
func NewAlibabajhscommunityidentitystoreRequest() *AlibabajhscommunityidentitystoreAPIRequest {
	return &AlibabajhscommunityidentitystoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajhscommunityidentitystoreAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.identity.store"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajhscommunityidentitystoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajhscommunityidentitystoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 会话token
func (r *AlibabajhscommunityidentitystoreAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabajhscommunityidentitystoreAPIRequest) GetToken() string {
	return r._token
}

// SetNickName is NickName Setter
// 用户昵称
func (r *AlibabajhscommunityidentitystoreAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// GetNickName NickName Getter
func (r AlibabajhscommunityidentitystoreAPIRequest) GetNickName() string {
	return r._nickName
}

// SetAvatarUrl is AvatarUrl Setter
// 用户头像
func (r *AlibabajhscommunityidentitystoreAPIRequest) SetAvatarUrl(_avatarUrl string) error {
	r._avatarUrl = _avatarUrl
	r.Set("avatar_url", _avatarUrl)
	return nil
}

// GetAvatarUrl AvatarUrl Getter
func (r AlibabajhscommunityidentitystoreAPIRequest) GetAvatarUrl() string {
	return r._avatarUrl
}
