package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSaleforcastPmMonthUpload 24-销售月预测数量（产管）回传-月度
// alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload
//
// 销售月预测数量（产管）回传-月度
func AlibabaTmallgenieScpPlanSaleforcastPmMonthUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
