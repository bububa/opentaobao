package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpSubmitAPIRequest
聚安全实人认证提交认证接口 API请求
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口 */
type AlibabaSecurityJaqRpSubmitAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
}

// New
