package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商家产品服务-编辑产品 
alibaba.wdk.merchantproduct.edit

商家产品服务-编辑产品
*/
func AlibabaWdkMerchantproductEdit(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantproductEditRequest, session string) (*wdk.AlibabaWdkMerchantproductEditAPIResponse, error) {
    var resp wdk.AlibabaWdkMerchantproductEditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
