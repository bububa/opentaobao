package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
二级物料-LT内的POGAP数据回传 
alibaba.tmallgenie.scp.plan.rawpo.gap.return

二级物料-LT内的POGAP数据回传
*/
func AlibabaTmallgenieScpPlanRawpoGapReturn(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanRawpoGapReturnRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
