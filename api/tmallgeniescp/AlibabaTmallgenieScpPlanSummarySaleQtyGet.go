package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSummarySaleQtyGet 同步销售数据按照渠道类型汇总
// alibaba.tmallgenie.scp.plan.summary.sale.qty.get
//
// 同步销售数据按照渠道类型汇总
func AlibabaTmallgenieScpPlanSummarySaleQtyGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
