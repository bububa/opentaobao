package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rt回传采购价 APIRequest
alibaba.wdk.purchase.price

猫超共享库存项目-rt回传采购价
*/
type AlibabaWdkPurchasePriceRequest struct {
    model.Params

    // 入参
    wdkOpenPurchasePrice   *WdkOpenPurchasePrice 

}

func NewAlibabaWdkPurchasePriceRequest() *AlibabaWdkPurchasePriceRequest{
    return &AlibabaWdkPurchasePriceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkPurchasePriceRequest) GetApiMethodName() string {
    return "alibaba.wdk.purchase.price"
}

func (r AlibabaWdkPurchasePriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkPurchasePriceRequest) SetWdkOpenPurchasePrice(wdkOpenPurchasePrice *WdkOpenPurchasePrice) error {
    r.wdkOpenPurchasePrice = wdkOpenPurchasePrice
    r.Set("wdk_open_purchase_price", wdkOpenPurchasePrice)
    return nil
}

func (r AlibabaWdkPurchasePriceRequest) GetWdkOpenPurchasePrice() *WdkOpenPurchasePrice {
    return r.wdkOpenPurchasePrice
}

