package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品标记接口 APIResponse
alibaba.wdk.sku.feature

给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
*/
type AlibabaWdkSkuFeatureAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuFeatureResponse `json:"alibaba_wdk_sku_feature_response,omitempty"` 
    AlibabaWdkSkuFeatureResponse
}

/* model for simplify = false
type AlibabaWdkSkuFeatureResponse struct {

    // 根据站点名称查询产品
    
    Result  *struct {
        AlibabaWdkSkuFeatureApiResult  *AlibabaWdkSkuFeatureApiResult `json:"alibaba_wdk_sku_feature_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuFeatureResponse struct {

    // 根据站点名称查询产品
    Result   *AlibabaWdkSkuFeatureApiResult `json:"result,omitempty"`

}
