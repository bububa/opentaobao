package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanMaterialPurchaseAttrGet 物料的采购属性查询
// alibaba.tmallgenie.scp.plan.material.purchase.attr.get
//
// 物料的采购属性查询
func AlibabaTmallgenieScpPlanMaterialPurchaseAttrGet(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
