package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
11-同步本周的po单（从W-1周到W+4周） 
alibaba.tmallgenie.scp.plan.current.po.get

11-同步本周的po单（从W-1周到W+4周）
*/
func AlibabaTmallgenieScpPlanCurrentPoGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCurrentPoGetRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
