package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanMaterialPurchaseAttrGet 物料的采购属性查询
// alibaba.tmallgenie.scp.plan.material.purchase.attr.get
//
// 物料的采购属性查询
func AlibabaTmallgenieScpPlanMaterialPurchaseAttrGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
