package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 采购价上传接口 APIRequest
alibaba.newretail.purchase.price.save

共享库存业务 供应商上传商品采购价
*/
type AlibabaNewretailPurchasePriceSaveRequest struct {
    model.Params

    // 接口入参
    savePurchasePriceRequest   *SavePurchasePriceRequest 

}

func NewAlibabaNewretailPurchasePriceSaveRequest() *AlibabaNewretailPurchasePriceSaveRequest{
    return &AlibabaNewretailPurchasePriceSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNewretailPurchasePriceSaveRequest) GetApiMethodName() string {
    return "alibaba.newretail.purchase.price.save"
}

func (r AlibabaNewretailPurchasePriceSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNewretailPurchasePriceSaveRequest) SetSavePurchasePriceRequest(savePurchasePriceRequest *SavePurchasePriceRequest) error {
    r.savePurchasePriceRequest = savePurchasePriceRequest
    r.Set("save_purchase_price_request", savePurchasePriceRequest)
    return nil
}

func (r AlibabaNewretailPurchasePriceSaveRequest) GetSavePurchasePriceRequest() *SavePurchasePriceRequest {
    return r.savePurchasePriceRequest
}

