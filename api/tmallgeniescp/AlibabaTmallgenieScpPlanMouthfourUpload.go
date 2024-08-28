package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanMouthfourUpload 21-M+4PR 回传接口接口
// alibaba.tmallgenie.scp.plan.mouthfour.upload
//
// M+4 PR 回传接口
func AlibabaTmallgenieScpPlanMouthfourUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanMouthfourUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
