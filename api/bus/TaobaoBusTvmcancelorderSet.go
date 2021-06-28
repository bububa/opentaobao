package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
线下自助机未付款取消订单 
taobao.bus.tvmcancelorder.set

自助机汽车票未付款取消订单
*/
func TaobaoBusTvmcancelorderSet(clt *core.SDKClient, req *bus.TaobaoBusTvmcancelorderSetRequest, session string) (*bus.TaobaoBusTvmcancelorderSetAPIResponse, error) {
    var resp bus.TaobaoBusTvmcancelorderSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
