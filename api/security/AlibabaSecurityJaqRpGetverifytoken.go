package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpGetverifytoken 聚安全实人认证获取认证会话token
// alibaba.security.jaq.rp.getverifytoken
//
// 聚安全实人认证获取认证会话token
func AlibabaSecurityJaqRpGetverifytoken(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpGetverifytokenAPIRequest, resp *security.AlibabaSecurityJaqRpGetverifytokenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
