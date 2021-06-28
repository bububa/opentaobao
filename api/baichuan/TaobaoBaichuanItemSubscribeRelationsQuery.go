package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
按条件查询订阅关系 
taobao.baichuan.item.subscribe.relations.query

按条件查询订阅关系
*/
func TaobaoBaichuanItemSubscribeRelationsQuery(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeRelationsQueryRequest, session string) (*baichuan.TaobaoBaichuanItemSubscribeRelationsQueryAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanItemSubscribeRelationsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
