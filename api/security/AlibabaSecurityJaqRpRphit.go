package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpRphit 聚安全-实人认证日志打点接口
// alibaba.security.jaq.rp.rphit
//
// 聚安全实人认证日志打点接口
func AlibabaSecurityJaqRpRphit(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpRphitAPIRequest, resp *security.AlibabaSecurityJaqRpRphitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
