package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
查询租赁商品信息 
alibaba.idle.rent.item.query

查询租赁商品信息
*/
func AlibabaIdleRentItemQuery(clt *core.SDKClient, req *idle.AlibabaIdleRentItemQueryRequest, session string) (*idle.AlibabaIdleRentItemQueryAPIResponse, error) {
    var resp idle.AlibabaIdleRentItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
