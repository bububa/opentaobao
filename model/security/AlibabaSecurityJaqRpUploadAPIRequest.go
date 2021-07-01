package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpUploadAPIRequest
聚安全实人认证上传认证信息 API请求
alibaba.security.jaq.rp.upload

聚安全实人认证上传认证信息 */
type AlibabaSecurityJaqRpUploadAPIRequest struct {
	model.Params
	// 认证会话token
	_verifyToken string
	// 此次需要上传的认证信息的列表
	_elements []Element
}

// New
