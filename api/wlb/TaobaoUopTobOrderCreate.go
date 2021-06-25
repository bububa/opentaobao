package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
ToB仓储发货 
taobao.uop.tob.order.create

ToB仓储发货
*/
func TaobaoUopTobOrderCreate(clt *core.SDKClient, req *wlb.TaobaoUopTobOrderCreateRequest, session string) (*wlb.TaobaoUopTobOrderCreateResponse, error) {
    var resp wlb.TaobaoUopTobOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
