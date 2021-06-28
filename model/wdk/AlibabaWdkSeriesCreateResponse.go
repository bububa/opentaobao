package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-新增系列 APIResponse
alibaba.wdk.series.create

系列品变更-新增系列
*/
type AlibabaWdkSeriesCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSeriesCreateResponse `json:"alibaba_wdk_series_create_response,omitempty"` 
    AlibabaWdkSeriesCreateResponse
}

/* model for simplify = false
type AlibabaWdkSeriesCreateResponse struct {

    // 调用结果
    
    ApiResult  *struct {
        AlibabaWdkSeriesCreateApiResult  *AlibabaWdkSeriesCreateApiResult `json:"alibaba_wdk_series_create_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkSeriesCreateResponse struct {

    // 调用结果
    ApiResult   *AlibabaWdkSeriesCreateApiResult `json:"api_result,omitempty"`

}
