package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudStart 实人认证云开始认证
// alibaba.security.jaq.rp.cloud.start
//
// 聚安全实人认证开始
func AlibabaSecurityJaqRpCloudStart(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudStartAPIRequest, resp *security.AlibabaSecurityJaqRpCloudStartAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
