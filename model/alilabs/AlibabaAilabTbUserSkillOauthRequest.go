package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户技能 Oauth 授权（淘宝 openId） APIRequest
alibaba.ailab.tb.user.skill.oauth

定制机厂商，在用户配网完成后，厂商调用此接口，写入特定技能的 Oauth 信息
*/
type AlibabaAilabTbUserSkillOauthRequest struct {
    model.Params

    // taobao open id
    taobaoId   string 

    // access token 过期时间，单位：秒
    expireIn   int64 

    // access token
    oauthAccessToken   string 

    // refresh token
    refreshToken   string 

}

func NewAlibabaAilabTbUserSkillOauthRequest() *AlibabaAilabTbUserSkillOauthRequest{
    return &AlibabaAilabTbUserSkillOauthRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabTbUserSkillOauthRequest) GetApiMethodName() string {
    return "alibaba.ailab.tb.user.skill.oauth"
}

func (r AlibabaAilabTbUserSkillOauthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabTbUserSkillOauthRequest) SetTaobaoId(taobaoId string) error {
    r.taobaoId = taobaoId
    r.Set("taobao_id", taobaoId)
    return nil
}

func (r AlibabaAilabTbUserSkillOauthRequest) GetTaobaoId() string {
    return r.taobaoId
}

func (r *AlibabaAilabTbUserSkillOauthRequest) SetExpireIn(expireIn int64) error {
    r.expireIn = expireIn
    r.Set("expire_in", expireIn)
    return nil
}

func (r AlibabaAilabTbUserSkillOauthRequest) GetExpireIn() int64 {
    return r.expireIn
}

func (r *AlibabaAilabTbUserSkillOauthRequest) SetOauthAccessToken(oauthAccessToken string) error {
    r.oauthAccessToken = oauthAccessToken
    r.Set("oauth_access_token", oauthAccessToken)
    return nil
}

func (r AlibabaAilabTbUserSkillOauthRequest) GetOauthAccessToken() string {
    return r.oauthAccessToken
}

func (r *AlibabaAilabTbUserSkillOauthRequest) SetRefreshToken(refreshToken string) error {
    r.refreshToken = refreshToken
    r.Set("refresh_token", refreshToken)
    return nil
}

func (r AlibabaAilabTbUserSkillOauthRequest) GetRefreshToken() string {
    return r.refreshToken
}

