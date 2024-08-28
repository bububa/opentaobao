package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSaleforcastPmUpload 18-销售预测数量（产管）回传接口
// alibaba.tmallgenie.scp.plan.saleforcast.pm.upload
//
// 销售预测数量（产管）回传接口
func AlibabaTmallgenieScpPlanSaleforcastPmUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
