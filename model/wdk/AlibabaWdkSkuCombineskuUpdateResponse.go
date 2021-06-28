package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合商品更新接口 APIResponse
alibaba.wdk.sku.combinesku.update

组合商品修改接口
*/
type AlibabaWdkSkuCombineskuUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuCombineskuUpdateResponse `json:"alibaba_wdk_sku_combinesku_update_response,omitempty"` 
    AlibabaWdkSkuCombineskuUpdateResponse
}

/* model for simplify = false
type AlibabaWdkSkuCombineskuUpdateResponse struct {

    // 调用结果
    
    Result  *struct {
        AlibabaWdkSkuCombineskuUpdateApiResults  *AlibabaWdkSkuCombineskuUpdateApiResults `json:"alibaba_wdk_sku_combinesku_update_api_results,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuCombineskuUpdateResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCombineskuUpdateApiResults `json:"result,omitempty"`

}
