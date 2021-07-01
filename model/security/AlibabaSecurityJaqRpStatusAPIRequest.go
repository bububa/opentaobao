package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpStatusAPIRequest
聚安全实人认证查询状态接口 API请求
alibaba.security.jaq.rp.status

聚安全实人认证查询状态接口 */
type AlibabaSecurityJaqRpStatusAPIRequest struct {
	model.Params
	// 账号id
	_accountId string
	// 凭据id
	_ticketId string
	// 客户端来源
	_source string
	// 业务来源
	_biz string
}

// New
