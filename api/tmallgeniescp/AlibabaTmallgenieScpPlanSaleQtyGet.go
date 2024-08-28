package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSaleQtyGet 12-同步销售数据
// alibaba.tmallgenie.scp.plan.sale.qty.get
//
// 同步销售数据
func AlibabaTmallgenieScpPlanSaleQtyGet(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
