package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqSpamregisterpreventionResultFetchNew 获取虚假注册保护结果
// alibaba.security.jaq.spamregisterprevention.result.fetch.new
//
// 获取虚假注册保护结果
func AlibabaSecurityJaqSpamregisterpreventionResultFetchNew(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIRequest, resp *security.AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
