package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
获取商品 
alibaba.data.item.get

获取商品信息，作为客户端Weex鉴权的虚拟api
*/
func AlibabaDataItemGet(clt *core.SDKClient, req *shop.AlibabaDataItemGetRequest, session string) (*shop.AlibabaDataItemGetResponse, error) {
    var resp shop.AlibabaDataItemGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
