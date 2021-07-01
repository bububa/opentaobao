package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* AlibabaTmallgenieScpPlanHistorySaleQtyGet
【已废除】同步历史的销售数据
alibaba.tmallgenie.scp.plan.history.sale.qty.get

同步历史的销售数据 */
func AlibabaTmallgenieScpPlanHistorySaleQtyGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
