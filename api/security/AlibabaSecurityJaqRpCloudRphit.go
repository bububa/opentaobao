package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudRphit 实人认证云服务日志打点
// alibaba.security.jaq.rp.cloud.rphit
//
// 聚安全实人认证日志打点接口
func AlibabaSecurityJaqRpCloudRphit(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudRphitAPIRequest, resp *security.AlibabaSecurityJaqRpCloudRphitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
