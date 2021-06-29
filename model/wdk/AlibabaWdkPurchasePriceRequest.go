package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rt回传采购价 API请求
alibaba.wdk.purchase.price

猫超共享库存项目-rt回传采购价
*/
type AlibabaWdkPurchasePriceRequest struct {
    model.Params
    // 入参
    _wdkOpenPurchasePrice   *WdkOpenPurchasePrice
}

// 初始化AlibabaWdkPurchasePriceRequest对象
func NewAlibabaWdkPurchasePriceRequest() *AlibabaWdkPurchasePriceRequest{
    return &AlibabaWdkPurchasePriceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPurchasePriceRequest) GetApiMethodName() string {
    return "alibaba.wdk.purchase.price"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPurchasePriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WdkOpenPurchasePrice Setter
// 入参
func (r *AlibabaWdkPurchasePriceRequest) SetWdkOpenPurchasePrice(_wdkOpenPurchasePrice *WdkOpenPurchasePrice) error {
    r._wdkOpenPurchasePrice = _wdkOpenPurchasePrice
    r.Set("wdk_open_purchase_price", _wdkOpenPurchasePrice)
    return nil
}

// WdkOpenPurchasePrice Getter
func (r AlibabaWdkPurchasePriceRequest) GetWdkOpenPurchasePrice() *WdkOpenPurchasePrice {
    return r._wdkOpenPurchasePrice
}
