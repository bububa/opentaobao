package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川批量商品订阅 
taobao.baichuan.items.subscribe

百川批量添加订阅的商品
*/
func TaobaoBaichuanItemsSubscribe(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsSubscribeRequest, session string) (*baichuan.TaobaoBaichuanItemsSubscribeResponse, error) {
    var resp baichuan.TaobaoBaichuanItemsSubscribeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
