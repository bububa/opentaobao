package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
新建门店商品 
alibaba.wdk.merchant.storeitem.create

新建门店商品
*/
func AlibabaWdkMerchantStoreitemCreate(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantStoreitemCreateRequest, session string) (*wdk.AlibabaWdkMerchantStoreitemCreateAPIResponse, error) {
    var resp wdk.AlibabaWdkMerchantStoreitemCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
