package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqSpamregisterpreventionResultFetch 获取垃圾注册防控结果
// alibaba.security.jaq.spamregisterprevention.result.fetch
//
// 获取垃圾注册防控结果
func AlibabaSecurityJaqSpamregisterpreventionResultFetch(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest, resp *security.AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
