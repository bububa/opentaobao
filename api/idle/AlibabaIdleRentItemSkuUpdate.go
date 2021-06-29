package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
更新/增加sku信息 
alibaba.idle.rent.item.sku.update

更新/增加sku信息
*/
func AlibabaIdleRentItemSkuUpdate(clt *core.SDKClient, req *idle.AlibabaIdleRentItemSkuUpdateRequest, session string) (*idle.AlibabaIdleRentItemSkuUpdateAPIResponse, error) {
    var resp idle.AlibabaIdleRentItemSkuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
