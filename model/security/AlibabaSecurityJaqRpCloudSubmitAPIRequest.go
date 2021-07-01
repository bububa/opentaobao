package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpCloudSubmitAPIRequest
实人认证云服务提交接口 API请求
alibaba.security.jaq.rp.cloud.submit

聚安全实人认证提交认证接口 */
type AlibabaSecurityJaqRpCloudSubmitAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
}

// New
