package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSummarySaleQtyGet 同步销售数据按照渠道类型汇总
// alibaba.tmallgenie.scp.plan.summary.sale.qty.get
//
// 同步销售数据按照渠道类型汇总
func AlibabaTmallgenieScpPlanSummarySaleQtyGet(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
