package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
汽车票车次查询 
taobao.bus.busnumber.get

提供汽车票车次查询服务
*/
func TaobaoBusBusnumberGet(clt *core.SDKClient, req *bus.TaobaoBusBusnumberGetRequest, session string) (*bus.TaobaoBusBusnumberGetResponse, error) {
    var resp bus.TaobaoBusBusnumberGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
