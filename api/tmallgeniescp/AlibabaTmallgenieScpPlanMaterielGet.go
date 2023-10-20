package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanMaterielGet 1-IBP同步物料接口
// alibaba.tmallgenie.scp.plan.materiel.get
//
// IBP同步物料接口
func AlibabaTmallgenieScpPlanMaterielGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanMaterielGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanMaterielGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
