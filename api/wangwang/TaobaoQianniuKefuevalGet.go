package wangwang

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wangwang"
)

/* 
客服评价详情接口 
taobao.qianniu.kefueval.get

获取买家对客服的服务评价
*/
func TaobaoQianniuKefuevalGet(clt *core.SDKClient, req *wangwang.TaobaoQianniuKefuevalGetRequest, session string) (*wangwang.TaobaoQianniuKefuevalGetAPIResponse, error) {
    var resp wangwang.TaobaoQianniuKefuevalGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
