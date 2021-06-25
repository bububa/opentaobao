package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
取消订单 
taobao.bus.cancleorder.set

取消订单
*/
func TaobaoBusCancleorderSet(clt *core.SDKClient, req *bus.TaobaoBusCancleorderSetRequest, session string) (*bus.TaobaoBusCancleorderSetResponse, error) {
    var resp bus.TaobaoBusCancleorderSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
