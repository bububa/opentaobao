package tmallhk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallhk"
)

/* 
天猫国际订单清关信息 
tmall.hk.order.clearance.query

天猫国际订单清关信息查询
*/
func TmallHkOrderClearanceQuery(clt *core.SDKClient, req *tmallhk.TmallHkOrderClearanceQueryRequest, session string) (*tmallhk.TmallHkOrderClearanceQueryAPIResponse, error) {
    var resp tmallhk.TmallHkOrderClearanceQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
