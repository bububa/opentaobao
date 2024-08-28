package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanRawpoGapReturn 二级物料-LT内的POGAP数据回传
// alibaba.tmallgenie.scp.plan.rawpo.gap.return
//
// 二级物料-LT内的POGAP数据回传
func AlibabaTmallgenieScpPlanRawpoGapReturn(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
