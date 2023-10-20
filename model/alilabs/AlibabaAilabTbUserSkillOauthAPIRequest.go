package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabTbUserSkillOauthAPIRequest 用户技能 Oauth 授权（淘宝 openId） API请求
// alibaba.ailab.tb.user.skill.oauth
//
// 定制机厂商，在用户配网完成后，厂商调用此接口，写入特定技能的 Oauth 信息
type AlibabaAilabTbUserSkillOauthAPIRequest struct {
	model.Params
	// taobao open id
	_taobaoId string
	// access token
	_oauthAccessToken string
	// refresh token
	_refreshToken string
	// access token 过期时间，单位：秒
	_expireIn int64
}

// NewAlibabaAilabTbUserSkillOauthRequest 初始化AlibabaAilabTbUserSkillOauthAPIRequest对象
func NewAlibabaAilabTbUserSkillOauthRequest() *AlibabaAilabTbUserSkillOauthAPIRequest {
	return &AlibabaAilabTbUserSkillOauthAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabTbUserSkillOauthAPIRequest) Reset() {
	r._taobaoId = ""
	r._oauthAccessToken = ""
	r._refreshToken = ""
	r._expireIn = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabTbUserSkillOauthAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.tb.user.skill.oauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabTbUserSkillOauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabTbUserSkillOauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaobaoId is TaobaoId Setter
// taobao open id
func (r *AlibabaAilabTbUserSkillOauthAPIRequest) SetTaobaoId(_taobaoId string) error {
	r._taobaoId = _taobaoId
	r.Set("taobao_id", _taobaoId)
	return nil
}

// GetTaobaoId TaobaoId Getter
func (r AlibabaAilabTbUserSkillOauthAPIRequest) GetTaobaoId() string {
	return r._taobaoId
}

// SetOauthAccessToken is OauthAccessToken Setter
// access token
func (r *AlibabaAilabTbUserSkillOauthAPIRequest) SetOauthAccessToken(_oauthAccessToken string) error {
	r._oauthAccessToken = _oauthAccessToken
	r.Set("oauth_access_token", _oauthAccessToken)
	return nil
}

// GetOauthAccessToken OauthAccessToken Getter
func (r AlibabaAilabTbUserSkillOauthAPIRequest) GetOauthAccessToken() string {
	return r._oauthAccessToken
}

// SetRefreshToken is RefreshToken Setter
// refresh token
func (r *AlibabaAilabTbUserSkillOauthAPIRequest) SetRefreshToken(_refreshToken string) error {
	r._refreshToken = _refreshToken
	r.Set("refresh_token", _refreshToken)
	return nil
}

// GetRefreshToken RefreshToken Getter
func (r AlibabaAilabTbUserSkillOauthAPIRequest) GetRefreshToken() string {
	return r._refreshToken
}

// SetExpireIn is ExpireIn Setter
// access token 过期时间，单位：秒
func (r *AlibabaAilabTbUserSkillOauthAPIRequest) SetExpireIn(_expireIn int64) error {
	r._expireIn = _expireIn
	r.Set("expire_in", _expireIn)
	return nil
}

// GetExpireIn ExpireIn Getter
func (r AlibabaAilabTbUserSkillOauthAPIRequest) GetExpireIn() int64 {
	return r._expireIn
}

var poolAlibabaAilabTbUserSkillOauthAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabTbUserSkillOauthRequest()
	},
}

// GetAlibabaAilabTbUserSkillOauthRequest 从 sync.Pool 获取 AlibabaAilabTbUserSkillOauthAPIRequest
func GetAlibabaAilabTbUserSkillOauthAPIRequest() *AlibabaAilabTbUserSkillOauthAPIRequest {
	return poolAlibabaAilabTbUserSkillOauthAPIRequest.Get().(*AlibabaAilabTbUserSkillOauthAPIRequest)
}

// ReleaseAlibabaAilabTbUserSkillOauthAPIRequest 将 AlibabaAilabTbUserSkillOauthAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabTbUserSkillOauthAPIRequest(v *AlibabaAilabTbUserSkillOauthAPIRequest) {
	v.Reset()
	poolAlibabaAilabTbUserSkillOauthAPIRequest.Put(v)
}
