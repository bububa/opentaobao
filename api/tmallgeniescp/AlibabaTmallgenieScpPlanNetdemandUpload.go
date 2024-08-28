package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanNetdemandUpload 23-Net Demand净需求回传
// alibaba.tmallgenie.scp.plan.netdemand.upload
//
// Net Demand净需求回传
func AlibabaTmallgenieScpPlanNetdemandUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
