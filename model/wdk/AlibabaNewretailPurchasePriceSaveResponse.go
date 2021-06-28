package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 采购价上传接口 APIResponse
alibaba.newretail.purchase.price.save

共享库存业务 供应商上传商品采购价
*/
type AlibabaNewretailPurchasePriceSaveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaNewretailPurchasePriceSaveResponse `json:"alibaba_newretail_purchase_price_save_response,omitempty"` 
    AlibabaNewretailPurchasePriceSaveResponse
}

/* model for simplify = false
type AlibabaNewretailPurchasePriceSaveResponse struct {

    // 调用结果对象
    
    Result  *struct {
        TopBaseResult  *TopBaseResult `json:"top_base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaNewretailPurchasePriceSaveResponse struct {

    // 调用结果对象
    Result   *TopBaseResult `json:"result,omitempty"`

}
