package idleisv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleisv"
)

/* 
服务商闲鱼商品发布 
alibaba.idle.isv.item.publish

服务商ISV闲鱼商品发布
*/
func AlibabaIdleIsvItemPublish(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemPublishAPIRequest, session string) (*idleisv.AlibabaIdleIsvItemPublishAPIResponse, error) {
    var resp idleisv.AlibabaIdleIsvItemPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
