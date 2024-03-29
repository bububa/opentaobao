package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudUpload 实人认证云上传接口
// alibaba.security.jaq.rp.cloud.upload
//
// 聚安全实人认证上传认证信息
func AlibabaSecurityJaqRpCloudUpload(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudUploadAPIRequest, resp *security.AlibabaSecurityJaqRpCloudUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
