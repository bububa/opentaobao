package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
二级物料-PO数据同步 
alibaba.tmallgenie.scp.plan.current.rawpo.get

二级物料-PO数据同步（WO-W[TL])
*/
func AlibabaTmallgenieScpPlanCurrentRawpoGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCurrentRawpoGetRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
