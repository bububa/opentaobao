package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanConsensusDemandUpload 20-IBP共识需求回传接口
// alibaba.tmallgenie.scp.plan.consensus.demand.upload
//
// IBP共识需求回传接口
func AlibabaTmallgenieScpPlanConsensusDemandUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanConsensusDemandUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
