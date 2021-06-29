package idleisv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleisv"
)

/* 
服务商闲鱼商品查询 
alibaba.idle.isv.item.query

服务商ISV闲鱼商品查询
*/
func AlibabaIdleIsvItemQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemQueryRequest, session string) (*idleisv.AlibabaIdleIsvItemQueryAPIResponse, error) {
    var resp idleisv.AlibabaIdleIsvItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
