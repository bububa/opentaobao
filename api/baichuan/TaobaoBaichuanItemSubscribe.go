package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
单个商品订阅 
taobao.baichuan.item.subscribe

百川单个商品订阅
*/
func TaobaoBaichuanItemSubscribe(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeRequest, session string) (*baichuan.TaobaoBaichuanItemSubscribeAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanItemSubscribeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
