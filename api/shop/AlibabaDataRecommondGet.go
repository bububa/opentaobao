package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
获取推荐信息 
alibaba.data.recommond.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
func AlibabaDataRecommondGet(clt *core.SDKClient, req *shop.AlibabaDataRecommondGetRequest, session string) (*shop.AlibabaDataRecommondGetResponse, error) {
    var resp shop.AlibabaDataRecommondGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
