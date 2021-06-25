package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
修改门店商品 
alibaba.wdk.merchant.storeitem.update

修改门店商品
*/
func AlibabaWdkMerchantStoreitemUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantStoreitemUpdateRequest, session string) (*wdk.AlibabaWdkMerchantStoreitemUpdateResponse, error) {
    var resp wdk.AlibabaWdkMerchantStoreitemUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
