package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系列品-商品排序 APIResponse
alibaba.wdk.series.sort

系列品商品变更-商品排序
*/
type AlibabaWdkSeriesSortAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSeriesSortResponse `json:"alibaba_wdk_series_sort_response,omitempty"` 
    AlibabaWdkSeriesSortResponse
}

/* model for simplify = false
type AlibabaWdkSeriesSortResponse struct {

    // 调用结果
    
    ApiResult  *struct {
        AlibabaWdkSeriesSortApiResult  *AlibabaWdkSeriesSortApiResult `json:"alibaba_wdk_series_sort_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkSeriesSortResponse struct {

    // 调用结果
    ApiResult   *AlibabaWdkSeriesSortApiResult `json:"api_result,omitempty"`

}
