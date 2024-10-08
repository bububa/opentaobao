package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanBomUpload 计划BOM同步
// alibaba.tmallgenie.scp.plan.bom.upload
//
// 计划BOM同步
func AlibabaTmallgenieScpPlanBomUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanBomUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanBomUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
