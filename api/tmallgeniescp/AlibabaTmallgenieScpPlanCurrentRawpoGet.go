package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanCurrentRawpoGet 二级物料-PO数据同步
// alibaba.tmallgenie.scp.plan.current.rawpo.get
//
// 二级物料-PO数据同步（WO-W[TL])
func AlibabaTmallgenieScpPlanCurrentRawpoGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCurrentRawpoGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
