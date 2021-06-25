package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
根据条件删除订阅关系 
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系
*/
func TaobaoBaichuanItemsUnsubscribeByCondition(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsUnsubscribeByConditionRequest, session string) (*baichuan.TaobaoBaichuanItemsUnsubscribeByConditionResponse, error) {
    var resp baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
