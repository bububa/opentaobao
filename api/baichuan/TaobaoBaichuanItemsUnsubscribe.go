package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
批量删除商品订阅 
taobao.baichuan.items.unsubscribe

批量删除商品订阅
*/
func TaobaoBaichuanItemsUnsubscribe(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsUnsubscribeRequest, session string) (*baichuan.TaobaoBaichuanItemsUnsubscribeAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanItemsUnsubscribeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
