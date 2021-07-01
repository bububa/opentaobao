package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAiUserQuickTokenBindAPIRequest
人工智能实验室精灵用户绑定第三方Token接口 API请求
alibaba.ai.user.quick.token.bind

人工智能实验室精灵用户绑定第三方Token接口 */
type AlibabaAiUserQuickTokenBindAPIRequest struct {
	model.Params
	// Oauth协议访问令牌
	_accessTokenValue string
	// 令牌有效期（秒）
	_expiredTime int64
	// 第三方用户账号唯一ID
	_merchantUserId string
	// 技能ID
	_skillId string
	// 账号隔离属性
	_schemaKey string
	// Oauth协议刷新令牌
	_refreshToken string
}

// New
