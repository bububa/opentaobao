package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
自助机条形码被动支付 
taobao.bus.tvmpayorder.set

汽车票线下自助机条形码支付
*/
func TaobaoBusTvmpayorderSet(clt *core.SDKClient, req *bus.TaobaoBusTvmpayorderSetRequest, session string) (*bus.TaobaoBusTvmpayorderSetResponse, error) {
    var resp bus.TaobaoBusTvmpayorderSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
