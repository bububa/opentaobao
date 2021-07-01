package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
门店查询接口 
alibaba.wdk.shop.query

根据门店code查询门店信息
*/
func AlibabaWdkShopQuery(clt *core.SDKClient, req *wdk.AlibabaWdkShopQueryAPIRequest, session string) (*wdk.AlibabaWdkShopQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkShopQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
