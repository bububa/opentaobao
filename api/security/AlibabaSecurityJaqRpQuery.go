package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpQuery 聚安全实人认证查询认证结果
// alibaba.security.jaq.rp.query
//
// 聚安全实人认证查询认证结果
func AlibabaSecurityJaqRpQuery(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpQueryAPIRequest, resp *security.AlibabaSecurityJaqRpQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
