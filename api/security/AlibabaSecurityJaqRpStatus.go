package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpStatus 聚安全实人认证查询状态接口
// alibaba.security.jaq.rp.status
//
// 聚安全实人认证查询状态接口
func AlibabaSecurityJaqRpStatus(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpStatusAPIRequest, resp *security.AlibabaSecurityJaqRpStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
