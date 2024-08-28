package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanChannelGet 5-IBP同步渠道接口
// alibaba.tmallgenie.scp.plan.channel.get
//
// IBP同步渠道接口
func AlibabaTmallgenieScpPlanChannelGet(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanChannelGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanChannelGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
