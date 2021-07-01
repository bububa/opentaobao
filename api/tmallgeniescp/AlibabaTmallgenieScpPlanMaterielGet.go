package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* AlibabaTmallgenieScpPlanMaterielGet
1-IBP同步物料接口
alibaba.tmallgenie.scp.plan.materiel.get

IBP同步物料接口 */
func AlibabaTmallgenieScpPlanMaterielGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanMaterielGetAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanMaterielGetAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanMaterielGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
