package idleisv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleisv"
)

/* 
服务商闲鱼商品编辑 
alibaba.idle.isv.item.edit

服务商ISV闲鱼商品编辑操作
*/
func AlibabaIdleIsvItemEdit(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemEditRequest, session string) (*idleisv.AlibabaIdleIsvItemEditAPIResponse, error) {
    var resp idleisv.AlibabaIdleIsvItemEditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
