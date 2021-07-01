package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
添加单个物流宝商品 
taobao.wlb.item.add

添加物流宝商品，支持物流宝子商品和属性添加
*/
func TaobaoWlbItemAdd(clt *core.SDKClient, req *wlb.TaobaoWlbItemAddAPIRequest, session string) (*wlb.TaobaoWlbItemAddAPIResponse, error) {
    var resp wlb.TaobaoWlbItemAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
