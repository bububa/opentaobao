package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpStatus 聚安全实人认证查询状态接口
// alibaba.security.jaq.rp.status
//
// 聚安全实人认证查询状态接口
func AlibabaSecurityJaqRpStatus(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpStatusAPIRequest, resp *security.AlibabaSecurityJaqRpStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
