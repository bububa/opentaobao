package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
物料的采购属性查询 
alibaba.tmallgenie.scp.plan.material.purchase.attr.get

物料的采购属性查询
*/
func AlibabaTmallgenieScpPlanMaterialPurchaseAttrGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanMaterialPurchaseAttrGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
