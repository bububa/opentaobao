package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudEvent 事件上报
// alibaba.security.jaq.rp.cloud.event
//
// 事件上报接口
func AlibabaSecurityJaqRpCloudEvent(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudEventAPIRequest, resp *security.AlibabaSecurityJaqRpCloudEventAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
