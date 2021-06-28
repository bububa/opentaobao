package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
线下自助机查询订单信息 
taobao.bus.tvmqueryorder.get

查询订单详情
*/
func TaobaoBusTvmqueryorderGet(clt *core.SDKClient, req *bus.TaobaoBusTvmqueryorderGetRequest, session string) (*bus.TaobaoBusTvmqueryorderGetAPIResponse, error) {
    var resp bus.TaobaoBusTvmqueryorderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
