package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 商户删除采购价 API请求
alibaba.newretail.purchase.price.delete

共享库存 商户删除采购价
*/
type AlibabaNewretailPurchasePriceDeleteRequest struct {
    model.Params
    // 调用入参
    deletePurchasePriceRequest   *DeletePurchasePriceRequest
}

// 初始化AlibabaNewretailPurchasePriceDeleteRequest对象
func NewAlibabaNewretailPurchasePriceDeleteRequest() *AlibabaNewretailPurchasePriceDeleteRequest{
    return &AlibabaNewretailPurchasePriceDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNewretailPurchasePriceDeleteRequest) GetApiMethodName() string {
    return "alibaba.newretail.purchase.price.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNewretailPurchasePriceDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeletePurchasePriceRequest Setter
// 调用入参
func (r *AlibabaNewretailPurchasePriceDeleteRequest) SetDeletePurchasePriceRequest(deletePurchasePriceRequest *DeletePurchasePriceRequest) error {
    r.deletePurchasePriceRequest = deletePurchasePriceRequest
    r.Set("delete_purchase_price_request", deletePurchasePriceRequest)
    return nil
}

// DeletePurchasePriceRequest Getter
func (r AlibabaNewretailPurchasePriceDeleteRequest) GetDeletePurchasePriceRequest() *DeletePurchasePriceRequest {
    return r.deletePurchasePriceRequest
}
