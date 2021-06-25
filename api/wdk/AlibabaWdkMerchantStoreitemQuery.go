package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
门店商品信心查询 
alibaba.wdk.merchant.storeitem.query

门店商品信心查询
*/
func AlibabaWdkMerchantStoreitemQuery(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantStoreitemQueryRequest, session string) (*wdk.AlibabaWdkMerchantStoreitemQueryResponse, error) {
    var resp wdk.AlibabaWdkMerchantStoreitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
