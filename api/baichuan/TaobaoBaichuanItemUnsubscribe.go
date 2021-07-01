package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
单个删除订阅关系 
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系
*/
func TaobaoBaichuanItemUnsubscribe(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemUnsubscribeAPIRequest, session string) (*baichuan.TaobaoBaichuanItemUnsubscribeAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanItemUnsubscribeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
