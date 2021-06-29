package idleitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleitem"
)

/* 
免费送商品发送 
alibaba.idle.item.idlecoin.add

免费送商品发布
*/
func AlibabaIdleItemIdlecoinAdd(clt *core.SDKClient, req *idleitem.AlibabaIdleItemIdlecoinAddRequest, session string) (*idleitem.AlibabaIdleItemIdlecoinAddAPIResponse, error) {
    var resp idleitem.AlibabaIdleItemIdlecoinAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
