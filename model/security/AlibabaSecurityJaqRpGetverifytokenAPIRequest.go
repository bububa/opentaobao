package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpGetverifytokenAPIRequest
聚安全实人认证获取认证会话token API请求
alibaba.security.jaq.rp.getverifytoken

聚安全实人认证获取认证会话token */
type AlibabaSecurityJaqRpGetverifytokenAPIRequest struct {
	model.Params
	// 账号，强烈建议填写，区别用户的唯一标识
	_accountId string
	// 选填，作为一次验证的唯一标识，每次验证需更换。如果不是验证类型可不填
	_ticketId string
	// 客户端来源
	_source string
	// 业务点
	_biz string
	// 额外信息
	_extraData string
}

// New
