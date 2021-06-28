package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增商品 APIResponse
alibaba.wdk.sku.add

创建RT门店商品或DC商品
*/
type AlibabaWdkSkuAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuAddResponse `json:"alibaba_wdk_sku_add_response,omitempty"` 
    AlibabaWdkSkuAddResponse
}

/* model for simplify = false
type AlibabaWdkSkuAddResponse struct {

    // 调用结果
    
    Result  *struct {
        AlibabaWdkSkuAddApiResults  *AlibabaWdkSkuAddApiResults `json:"alibaba_wdk_sku_add_api_results,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuAddResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuAddApiResults `json:"result,omitempty"`

}
