package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpCloudStartAPIRequest
实人认证云开始认证 API请求
alibaba.security.jaq.rp.cloud.start

聚安全实人认证开始 */
type AlibabaSecurityJaqRpCloudStartAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 客户端信息，如果是服务端接入，里面的参数可为空
	_clientInfo *RpClientInfo
	// 扩展信息
	_extraData string
}

// New
