package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpRphit 聚安全-实人认证日志打点接口
// alibaba.security.jaq.rp.rphit
//
// 聚安全实人认证日志打点接口
func AlibabaSecurityJaqRpRphit(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpRphitAPIRequest, resp *security.AlibabaSecurityJaqRpRphitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
