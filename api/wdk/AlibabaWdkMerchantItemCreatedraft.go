package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
新建商品草稿 
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口
*/
func AlibabaWdkMerchantItemCreatedraft(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantItemCreatedraftRequest, session string) (*wdk.AlibabaWdkMerchantItemCreatedraftResponse, error) {
    var resp wdk.AlibabaWdkMerchantItemCreatedraftAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
