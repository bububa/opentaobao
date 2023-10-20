package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanInventorQtyGet 10-同步库存现有量
// alibaba.tmallgenie.scp.plan.inventor.qty.get
//
// 同步库存现有量
func AlibabaTmallgenieScpPlanInventorQtyGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
