package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心价格修改 APIResponse
alibaba.wdk.item.price.update

修改门店商品售价和会员价
*/
type AlibabaWdkItemPriceUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkItemPriceUpdateResponse `json:"alibaba_wdk_item_price_update_response,omitempty"` 
    AlibabaWdkItemPriceUpdateResponse
}

/* model for simplify = false
type AlibabaWdkItemPriceUpdateResponse struct {

    // hsfResult
    
    HsfResult  *struct {
        AlibabaWdkItemPriceUpdateResult  *AlibabaWdkItemPriceUpdateResult `json:"alibaba_wdk_item_price_update_result,omitempty"`
    } `json:"hsf_result,omitempty"`
    

}
*/

type AlibabaWdkItemPriceUpdateResponse struct {

    // hsfResult
    HsfResult   *AlibabaWdkItemPriceUpdateResult `json:"hsf_result,omitempty"`

}
