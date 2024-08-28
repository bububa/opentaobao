package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpUpload 聚安全实人认证上传认证信息
// alibaba.security.jaq.rp.upload
//
// 聚安全实人认证上传认证信息
func AlibabaSecurityJaqRpUpload(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpUploadAPIRequest, resp *security.AlibabaSecurityJaqRpUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
