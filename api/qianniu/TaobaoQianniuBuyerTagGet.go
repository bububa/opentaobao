package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
判断买家是否有某些标 
taobao.qianniu.buyer.tag.get

判断某个买家是否有某些标
*/
func TaobaoQianniuBuyerTagGet(clt *core.SDKClient, req *qianniu.TaobaoQianniuBuyerTagGetAPIRequest, session string) (*qianniu.TaobaoQianniuBuyerTagGetAPIResponse, error) {
    var resp qianniu.TaobaoQianniuBuyerTagGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
