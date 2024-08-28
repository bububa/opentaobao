package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpFetchmaterial 聚安全实人认证获取结果接口
// alibaba.security.jaq.rp.fetchmaterial
//
// 聚安全实人认证获取结果接口
func AlibabaSecurityJaqRpFetchmaterial(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpFetchmaterialAPIRequest, resp *security.AlibabaSecurityJaqRpFetchmaterialAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
