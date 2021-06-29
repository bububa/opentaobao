package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 采购价上传接口 API请求
alibaba.newretail.purchase.price.save

共享库存业务 供应商上传商品采购价
*/
type AlibabaNewretailPurchasePriceSaveRequest struct {
    model.Params
    // 接口入参
    savePurchasePriceRequest   *SavePurchasePriceRequest
}

// 初始化AlibabaNewretailPurchasePriceSaveRequest对象
func NewAlibabaNewretailPurchasePriceSaveRequest() *AlibabaNewretailPurchasePriceSaveRequest{
    return &AlibabaNewretailPurchasePriceSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNewretailPurchasePriceSaveRequest) GetApiMethodName() string {
    return "alibaba.newretail.purchase.price.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNewretailPurchasePriceSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SavePurchasePriceRequest Setter
// 接口入参
func (r *AlibabaNewretailPurchasePriceSaveRequest) SetSavePurchasePriceRequest(savePurchasePriceRequest *SavePurchasePriceRequest) error {
    r.savePurchasePriceRequest = savePurchasePriceRequest
    r.Set("save_purchase_price_request", savePurchasePriceRequest)
    return nil
}

// SavePurchasePriceRequest Getter
func (r AlibabaNewretailPurchasePriceSaveRequest) GetSavePurchasePriceRequest() *SavePurchasePriceRequest {
    return r.savePurchasePriceRequest
}
