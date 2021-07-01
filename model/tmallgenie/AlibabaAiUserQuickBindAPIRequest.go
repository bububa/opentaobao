package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAiUserQuickBindAPIRequest
精灵用户绑定第三方账号信息 API请求
alibaba.ai.user.quick.bind

人工智能实验室精灵用户绑定第三方账号信息接口，开放给Iot厂商做为厂商上送第三方账号信息的接口 */
type AlibabaAiUserQuickBindAPIRequest struct {
	model.Params
	// 交易流水号（唯一即可，不参与业务运算）
	_serialNo string
	// 第三方用户类型
	_extUserType string
	// 第三用户ID
	_extUserId string
	// 请求时间
	_reqTime string
	// 商户的用户的唯一ID
	_merchantUserId string
	// 开放平台申请的schema
	_schemaKey string
}

// New
