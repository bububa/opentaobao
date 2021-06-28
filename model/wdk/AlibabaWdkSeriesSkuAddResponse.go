package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-添加商品 APIResponse
alibaba.wdk.series.sku.add

系列品商品变更-添加商品
*/
type AlibabaWdkSeriesSkuAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSeriesSkuAddResponse `json:"alibaba_wdk_series_sku_add_response,omitempty"` 
    AlibabaWdkSeriesSkuAddResponse
}

/* model for simplify = false
type AlibabaWdkSeriesSkuAddResponse struct {

    // 调用结果
    
    ApiResult  *struct {
        AlibabaWdkSeriesSkuAddApiResult  *AlibabaWdkSeriesSkuAddApiResult `json:"alibaba_wdk_series_sku_add_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkSeriesSkuAddResponse struct {

    // 调用结果
    ApiResult   *AlibabaWdkSeriesSkuAddApiResult `json:"api_result,omitempty"`

}
