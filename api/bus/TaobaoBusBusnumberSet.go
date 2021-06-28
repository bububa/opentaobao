package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
商家汽车票车次更新通知接口 
taobao.bus.busnumber.set

商家汽车票车次更新后，调用该接口通知平台。
*/
func TaobaoBusBusnumberSet(clt *core.SDKClient, req *bus.TaobaoBusBusnumberSetRequest, session string) (*bus.TaobaoBusBusnumberSetAPIResponse, error) {
    var resp bus.TaobaoBusBusnumberSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
