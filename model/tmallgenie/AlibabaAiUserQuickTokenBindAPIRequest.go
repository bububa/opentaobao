package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiUserQuickTokenBindAPIRequest 人工智能实验室精灵用户绑定第三方Token接口 API请求
// alibaba.ai.user.quick.token.bind
//
// 人工智能实验室精灵用户绑定第三方Token接口
type AlibabaAiUserQuickTokenBindAPIRequest struct {
	model.Params
	// Oauth协议访问令牌
	_accessTokenValue string
	// 第三方用户账号唯一ID
	_merchantUserId string
	// 技能ID
	_skillId string
	// 账号隔离属性
	_schemaKey string
	// Oauth协议刷新令牌
	_refreshToken string
	// 令牌有效期（秒）
	_expiredTime int64
}

// NewAlibabaAiUserQuickTokenBindRequest 初始化AlibabaAiUserQuickTokenBindAPIRequest对象
func NewAlibabaAiUserQuickTokenBindRequest() *AlibabaAiUserQuickTokenBindAPIRequest {
	return &AlibabaAiUserQuickTokenBindAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAiUserQuickTokenBindAPIRequest) Reset() {
	r._accessTokenValue = ""
	r._merchantUserId = ""
	r._skillId = ""
	r._schemaKey = ""
	r._refreshToken = ""
	r._expiredTime = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.user.quick.token.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessTokenValue is AccessTokenValue Setter
// Oauth协议访问令牌
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetAccessTokenValue(_accessTokenValue string) error {
	r._accessTokenValue = _accessTokenValue
	r.Set("access_token_value", _accessTokenValue)
	return nil
}

// GetAccessTokenValue AccessTokenValue Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetAccessTokenValue() string {
	return r._accessTokenValue
}

// SetMerchantUserId is MerchantUserId Setter
// 第三方用户账号唯一ID
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// SetSkillId is SkillId Setter
// 技能ID
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetSkillId(_skillId string) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetSkillId() string {
	return r._skillId
}

// SetSchemaKey is SchemaKey Setter
// 账号隔离属性
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetRefreshToken is RefreshToken Setter
// Oauth协议刷新令牌
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetRefreshToken(_refreshToken string) error {
	r._refreshToken = _refreshToken
	r.Set("refresh_token", _refreshToken)
	return nil
}

// GetRefreshToken RefreshToken Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetRefreshToken() string {
	return r._refreshToken
}

// SetExpiredTime is ExpiredTime Setter
// 令牌有效期（秒）
func (r *AlibabaAiUserQuickTokenBindAPIRequest) SetExpiredTime(_expiredTime int64) error {
	r._expiredTime = _expiredTime
	r.Set("expired_time", _expiredTime)
	return nil
}

// GetExpiredTime ExpiredTime Getter
func (r AlibabaAiUserQuickTokenBindAPIRequest) GetExpiredTime() int64 {
	return r._expiredTime
}

var poolAlibabaAiUserQuickTokenBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAiUserQuickTokenBindRequest()
	},
}

// GetAlibabaAiUserQuickTokenBindRequest 从 sync.Pool 获取 AlibabaAiUserQuickTokenBindAPIRequest
func GetAlibabaAiUserQuickTokenBindAPIRequest() *AlibabaAiUserQuickTokenBindAPIRequest {
	return poolAlibabaAiUserQuickTokenBindAPIRequest.Get().(*AlibabaAiUserQuickTokenBindAPIRequest)
}

// ReleaseAlibabaAiUserQuickTokenBindAPIRequest 将 AlibabaAiUserQuickTokenBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAiUserQuickTokenBindAPIRequest(v *AlibabaAiUserQuickTokenBindAPIRequest) {
	v.Reset()
	poolAlibabaAiUserQuickTokenBindAPIRequest.Put(v)
}
