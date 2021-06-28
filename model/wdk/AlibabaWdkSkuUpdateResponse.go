package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新商品 APIResponse
alibaba.wdk.sku.update

开放商品更新接口
*/
type AlibabaWdkSkuUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuUpdateResponse `json:"alibaba_wdk_sku_update_response,omitempty"` 
    AlibabaWdkSkuUpdateResponse
}

/* model for simplify = false
type AlibabaWdkSkuUpdateResponse struct {

    // 执行结果
    
    Result  *struct {
        AlibabaWdkSkuUpdateApiResults  *AlibabaWdkSkuUpdateApiResults `json:"alibaba_wdk_sku_update_api_results,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuUpdateResponse struct {

    // 执行结果
    Result   *AlibabaWdkSkuUpdateApiResults `json:"result,omitempty"`

}
