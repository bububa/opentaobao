package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqLoginpreventionResultFetch 获取登录保护结果
// alibaba.security.jaq.loginprevention.result.fetch
//
// 获取登录保护结果
func AlibabaSecurityJaqLoginpreventionResultFetch(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest, resp *security.AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
