package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpStart 聚安全实人认证开始
// alibaba.security.jaq.rp.start
//
// 聚安全实人认证开始
func AlibabaSecurityJaqRpStart(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpStartAPIRequest, resp *security.AlibabaSecurityJaqRpStartAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
