package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人工智能实验室精灵用户绑定第三方Token接口 API请求
alibaba.ai.user.quick.token.bind

人工智能实验室精灵用户绑定第三方Token接口
*/
type AlibabaAiUserQuickTokenBindAPIRequest struct {
    model.Params
    // Oauth协议访问令牌
    _accessTokenValue   string
    // 令牌有效期（秒）
    _expiredTime   int64
    // 第三方用户账号唯一ID
    _merchantUserId   string
    // 技能ID
    _skillId   string
    // 账号隔离属性
    _schemaKey   string
    // Oauth协议刷新令牌
    _refreshToken   string
}

// 初始化AlibabaAiUserQuickTokenBindAPIRequest对象
func NewAlibabaAiUserQuickTokenBindRequest() *AlibabaAiUserQuickTokenBindAPIRequest{
    return &AlibabaAiUserQuickTokenBindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetApiMethodName() string {
    return "alibaba.ai.user.quick.token.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccessTokenValue Setter
// Oauth协议访问令牌
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetAccessTokenValue(_accessTokenValue string) error {
    r._accessTokenValue = _accessTokenValue
    r.Set("access_token_value", _accessTokenValue)
    return nil
}

// AccessTokenValue Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetAccessTokenValue() string {
    return r._accessTokenValue
}
// ExpiredTime Setter
// 令牌有效期（秒）
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetExpiredTime(_expiredTime int64) error {
    r._expiredTime = _expiredTime
    r.Set("expired_time", _expiredTime)
    return nil
}

// ExpiredTime Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetExpiredTime() int64 {
    return r._expiredTime
}
// MerchantUserId Setter
// 第三方用户账号唯一ID
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetMerchantUserId(_merchantUserId string) error {
    r._merchantUserId = _merchantUserId
    r.Set("merchant_user_id", _merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetMerchantUserId() string {
    return r._merchantUserId
}
// SkillId Setter
// 技能ID
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetSkillId(_skillId string) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetSkillId() string {
    return r._skillId
}
// SchemaKey Setter
// 账号隔离属性
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetSchemaKey() string {
    return r._schemaKey
}
// RefreshToken Setter
// Oauth协议刷新令牌
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetRefreshToken(_refreshToken string) error {
    r._refreshToken = _refreshToken
    r.Set("refresh_token", _refreshToken)
    return nil
}

// RefreshToken Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetRefreshToken() string {
    return r._refreshToken
}
