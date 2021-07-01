package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
共享库存 商户删除采购价 
alibaba.newretail.purchase.price.delete

共享库存 商户删除采购价
*/
func AlibabaNewretailPurchasePriceDelete(clt *core.SDKClient, req *wdk.AlibabaNewretailPurchasePriceDeleteAPIRequest, session string) (*wdk.AlibabaNewretailPurchasePriceDeleteAPIResponse, error) {
    var resp wdk.AlibabaNewretailPurchasePriceDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
