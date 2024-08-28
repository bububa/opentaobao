package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaMsfserviceAuditsUpdate 操作改约审批单
// alibaba.msfservice.audits.update
//
// 操作改约审批单
func AlibabaMsfserviceAuditsUpdate(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaMsfserviceAuditsUpdateAPIRequest, resp *tmallsc.AlibabaMsfserviceAuditsUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
