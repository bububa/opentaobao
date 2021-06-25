package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 商户删除采购价 APIRequest
alibaba.newretail.purchase.price.delete

共享库存 商户删除采购价
*/
type AlibabaNewretailPurchasePriceDeleteRequest struct {
    model.Params

    // 调用入参
    deletePurchasePriceRequest   *DeletePurchasePriceRequest 

}

func NewAlibabaNewretailPurchasePriceDeleteRequest() *AlibabaNewretailPurchasePriceDeleteRequest{
    return &AlibabaNewretailPurchasePriceDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNewretailPurchasePriceDeleteRequest) GetApiMethodName() string {
    return "alibaba.newretail.purchase.price.delete"
}

func (r AlibabaNewretailPurchasePriceDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNewretailPurchasePriceDeleteRequest) SetDeletePurchasePriceRequest(deletePurchasePriceRequest *DeletePurchasePriceRequest) error {
    r.deletePurchasePriceRequest = deletePurchasePriceRequest
    r.Set("delete_purchase_price_request", deletePurchasePriceRequest)
    return nil
}

func (r AlibabaNewretailPurchasePriceDeleteRequest) GetDeletePurchasePriceRequest() *DeletePurchasePriceRequest {
    return r.deletePurchasePriceRequest
}

