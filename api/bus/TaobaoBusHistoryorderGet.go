package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
历史订单查询（对账） 
taobao.bus.historyorder.get

历史订单查询，对账接口
*/
func TaobaoBusHistoryorderGet(clt *core.SDKClient, req *bus.TaobaoBusHistoryorderGetRequest, session string) (*bus.TaobaoBusHistoryorderGetAPIResponse, error) {
    var resp bus.TaobaoBusHistoryorderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
