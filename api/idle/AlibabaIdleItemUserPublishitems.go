package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
发布的商品列表 
alibaba.idle.item.user.publishitems

为服务商的卖家提供发布的闲鱼商品列表
*/
func AlibabaIdleItemUserPublishitems(clt *core.SDKClient, req *idle.AlibabaIdleItemUserPublishitemsRequest, session string) (*idle.AlibabaIdleItemUserPublishitemsAPIResponse, error) {
    var resp idle.AlibabaIdleItemUserPublishitemsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
