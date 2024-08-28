package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudSubmit 实人认证云服务提交接口
// alibaba.security.jaq.rp.cloud.submit
//
// 聚安全实人认证提交认证接口
func AlibabaSecurityJaqRpCloudSubmit(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudSubmitAPIRequest, resp *security.AlibabaSecurityJaqRpCloudSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
