package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
12-同步销售数据 
alibaba.tmallgenie.scp.plan.sale.qty.get

同步销售数据
*/
func AlibabaTmallgenieScpPlanSaleQtyGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
