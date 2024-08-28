package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpStart 聚安全实人认证开始
// alibaba.security.jaq.rp.start
//
// 聚安全实人认证开始
func AlibabaSecurityJaqRpStart(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpStartAPIRequest, resp *security.AlibabaSecurityJaqRpStartAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
