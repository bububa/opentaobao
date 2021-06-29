package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人工智能实验室精灵用户绑定第三方Token接口 APIRequest
alibaba.ai.user.quick.token.bind

人工智能实验室精灵用户绑定第三方Token接口
*/
type AlibabaAiUserQuickTokenBindRequest struct {
    model.Params

    // Oauth协议访问令牌
    accessTokenValue   string 

    // 令牌有效期（秒）
    expiredTime   int64 

    // 第三方用户账号唯一ID
    merchantUserId   string 

    // 技能ID
    skillId   string 

    // 账号隔离属性
    schemaKey   string 

    // Oauth协议刷新令牌
    refreshToken   string 

}

func NewAlibabaAiUserQuickTokenBindRequest() *AlibabaAiUserQuickTokenBindRequest{
    return &AlibabaAiUserQuickTokenBindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAiUserQuickTokenBindRequest) GetApiMethodName() string {
    return "alibaba.ai.user.quick.token.bind"
}

func (r AlibabaAiUserQuickTokenBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAiUserQuickTokenBindRequest) SetAccessTokenValue(accessTokenValue string) error {
    r.accessTokenValue = accessTokenValue
    r.Set("access_token_value", accessTokenValue)
    return nil
}

func (r AlibabaAiUserQuickTokenBindRequest) GetAccessTokenValue() string {
    return r.accessTokenValue
}

func (r *AlibabaAiUserQuickTokenBindRequest) SetExpiredTime(expiredTime int64) error {
    r.expiredTime = expiredTime
    r.Set("expired_time", expiredTime)
    return nil
}

func (r AlibabaAiUserQuickTokenBindRequest) GetExpiredTime() int64 {
    return r.expiredTime
}

func (r *AlibabaAiUserQuickTokenBindRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

func (r AlibabaAiUserQuickTokenBindRequest) GetMerchantUserId() string {
    return r.merchantUserId
}

func (r *AlibabaAiUserQuickTokenBindRequest) SetSkillId(skillId string) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r AlibabaAiUserQuickTokenBindRequest) GetSkillId() string {
    return r.skillId
}

func (r *AlibabaAiUserQuickTokenBindRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r AlibabaAiUserQuickTokenBindRequest) GetSchemaKey() string {
    return r.schemaKey
}

func (r *AlibabaAiUserQuickTokenBindRequest) SetRefreshToken(refreshToken string) error {
    r.refreshToken = refreshToken
    r.Set("refresh_token", refreshToken)
    return nil
}

func (r AlibabaAiUserQuickTokenBindRequest) GetRefreshToken() string {
    return r.refreshToken
}

