package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanCurrentPoGet 11-同步本周的po单（从W-1周到W+4周）
// alibaba.tmallgenie.scp.plan.current.po.get
//
// 11-同步本周的po单（从W-1周到W+4周）
func AlibabaTmallgenieScpPlanCurrentPoGet(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
