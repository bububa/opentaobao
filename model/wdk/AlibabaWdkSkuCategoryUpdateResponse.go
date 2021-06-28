package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家类目修改接口 APIResponse
alibaba.wdk.sku.category.update

商家类目修改接口
*/
type AlibabaWdkSkuCategoryUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuCategoryUpdateResponse `json:"alibaba_wdk_sku_category_update_response,omitempty"` 
    AlibabaWdkSkuCategoryUpdateResponse
}

/* model for simplify = false
type AlibabaWdkSkuCategoryUpdateResponse struct {

    // 调用结果
    
    Result  *struct {
        AlibabaWdkSkuCategoryUpdateApiResult  *AlibabaWdkSkuCategoryUpdateApiResult `json:"alibaba_wdk_sku_category_update_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuCategoryUpdateResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCategoryUpdateApiResult `json:"result,omitempty"`

}
