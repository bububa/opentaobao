package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAiUserQuickRegisterAPIRequest
精灵用户注册申请 API请求
alibaba.ai.user.quick.register

人工智能实验室精灵用户注册申请接口，开放给Iot厂商做厂商会员数据上报 */
type AlibabaAiUserQuickRegisterAPIRequest struct {
	model.Params
	// 请求交易流水号（唯一即可，不参与业务运算）
	_serialNo string
	// 请求时间
	_reqTime string
	// 商户的用户的唯一ID
	_merchantUserId string
	// 账户体系隔离
	_schemaKey string
}

// New
