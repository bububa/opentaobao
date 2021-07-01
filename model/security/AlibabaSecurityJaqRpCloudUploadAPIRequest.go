package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpCloudUploadAPIRequest
实人认证云上传接口 API请求
alibaba.security.jaq.rp.cloud.upload

聚安全实人认证上传认证信息 */
type AlibabaSecurityJaqRpCloudUploadAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
	// []
	_elements []Elements
}

// New
