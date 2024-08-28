package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpGetverifytoken 聚安全实人认证获取认证会话token
// alibaba.security.jaq.rp.getverifytoken
//
// 聚安全实人认证获取认证会话token
func AlibabaSecurityJaqRpGetverifytoken(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpGetverifytokenAPIRequest, resp *security.AlibabaSecurityJaqRpGetverifytokenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
