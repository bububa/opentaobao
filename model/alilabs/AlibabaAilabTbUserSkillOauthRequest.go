package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户技能 Oauth 授权（淘宝 openId） API请求
alibaba.ailab.tb.user.skill.oauth

定制机厂商，在用户配网完成后，厂商调用此接口，写入特定技能的 Oauth 信息
*/
type AlibabaAilabTbUserSkillOauthRequest struct {
    model.Params
    // taobao open id
    _taobaoId   string
    // access token 过期时间，单位：秒
    _expireIn   int64
    // access token
    _oauthAccessToken   string
    // refresh token
    _refreshToken   string
}

// 初始化AlibabaAilabTbUserSkillOauthRequest对象
func NewAlibabaAilabTbUserSkillOauthRequest() *AlibabaAilabTbUserSkillOauthRequest{
    return &AlibabaAilabTbUserSkillOauthRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabTbUserSkillOauthRequest) GetApiMethodName() string {
    return "alibaba.ailab.tb.user.skill.oauth"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabTbUserSkillOauthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaobaoId Setter
// taobao open id
func (r *AlibabaAilabTbUserSkillOauthRequest) SetTaobaoId(_taobaoId string) error {
    r._taobaoId = _taobaoId
    r.Set("taobao_id", _taobaoId)
    return nil
}

// TaobaoId Getter
func (r AlibabaAilabTbUserSkillOauthRequest) GetTaobaoId() string {
    return r._taobaoId
}
// ExpireIn Setter
// access token 过期时间，单位：秒
func (r *AlibabaAilabTbUserSkillOauthRequest) SetExpireIn(_expireIn int64) error {
    r._expireIn = _expireIn
    r.Set("expire_in", _expireIn)
    return nil
}

// ExpireIn Getter
func (r AlibabaAilabTbUserSkillOauthRequest) GetExpireIn() int64 {
    return r._expireIn
}
// OauthAccessToken Setter
// access token
func (r *AlibabaAilabTbUserSkillOauthRequest) SetOauthAccessToken(_oauthAccessToken string) error {
    r._oauthAccessToken = _oauthAccessToken
    r.Set("oauth_access_token", _oauthAccessToken)
    return nil
}

// OauthAccessToken Getter
func (r AlibabaAilabTbUserSkillOauthRequest) GetOauthAccessToken() string {
    return r._oauthAccessToken
}
// RefreshToken Setter
// refresh token
func (r *AlibabaAilabTbUserSkillOauthRequest) SetRefreshToken(_refreshToken string) error {
    r._refreshToken = _refreshToken
    r.Set("refresh_token", _refreshToken)
    return nil
}

// RefreshToken Getter
func (r AlibabaAilabTbUserSkillOauthRequest) GetRefreshToken() string {
    return r._refreshToken
}
