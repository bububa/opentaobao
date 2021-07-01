package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpQueryAPIRequest
聚安全实人认证查询认证结果 API请求
alibaba.security.jaq.rp.query

聚安全实人认证查询认证结果 */
type AlibabaSecurityJaqRpQueryAPIRequest struct {
	model.Params
	// token
	_verifyToken string
}

// New
