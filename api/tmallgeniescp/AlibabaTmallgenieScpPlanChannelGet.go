package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
5-IBP同步渠道接口 
alibaba.tmallgenie.scp.plan.channel.get

IBP同步渠道接口
*/
func AlibabaTmallgenieScpPlanChannelGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanChannelGetAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanChannelGetAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanChannelGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
